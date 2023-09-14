package pool

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func example1() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()
	instance := myPool.Get()

	myPool.Put(instance) // 2. 이전에 조회한 인스턴스를 풀에 돌려놓는다. 사용 가능한 인스턴스 수 1 증가
	myPool.Get()         // 3. 이전에 할당된 후 다시 풀에 넣은 인스턴스를 사용
}

func memoryExample() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem // 1. byte 슬라이스 들의 주소를 저장
		},
	}

	// 4KB 로 풀을 시작한다.
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte) // 2. 타입이 당연히 byte 슬라이스에 대한 포인터라고 가정
			defer calcPool.Put(mem)
			// 이 메모리에서는 뭔가 흥미롭지만 빠른 작업이 이루어진다고 가정
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}

func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}

			connectToService()

			fmt.Fprintln(conn, "")
			conn.Close()
		}
	}()

	return &wg
}

// sync.Pool 로 성능이 얼마나 개선되는지 확인
func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}

	for i := 0; i < 10; i++ {
		p.Put(p.New)
	}

	return p
}

func startNetworkDaemonUsingPool() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmServiceConnCache()

		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}

			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")

			connPool.Put(svcConn)
			conn.Close()
		}
	}()

	return &wg
}
