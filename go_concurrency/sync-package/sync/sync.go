package sync

import (
	"fmt"
	"sync"
	"time"
)

func cond() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()        // 8. 임계 영역 진입 했기 때문에 관련 데이터 수정 가능
		queue = queue[1:] // 9. 항목에서 큐 제거 역할
		fmt.Println("Removed from queue")
		c.L.Unlock() // 10. 항목 제거 후 임계 영역 벗어나기
		c.Signal()   // 11. 조건을 기다리는 고루틴에게 뭔가 발생했음을 전달
	}

	for i := 0; i < 10; i++ {
		c.L.Lock() // 3. 임계 영역 진입
		for len(queue) == 2 {
			c.Wait() // 5. 조건에 대한 신호가 전송될 때까지 main 고루틴은 일시 중단
		}

		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})

		go removeFromQueue(1 * time.Second) // 6. 1초 후에 항목을 큐에서 꺼내는 새로운 고루틴 생성
		c.L.Unlock()
	}
}

func broadcast() {
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	// 2. 조건들의 신호들을 처리하는 함수를 등록할 수 있는 편의 함수 정의
	// 고루틴이 실행 중이라는 것을 확인하기 전까지 subscribe 종료 X
	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)

		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()

		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)

	// 5. 버튼 클릭시 버튼 윈도우 최대화 시뮬레이션
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window")
		clickRegistered.Done()
	})

	// 6. 마우스 클릭 시 대화 상자 표시 사뮬레이션
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	// 7. 사용자 버튼 클릭 시뮬레이션
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast()

	clickRegistered.Wait()
}
