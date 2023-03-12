package once

import (
	"fmt"
	"sync"
)

func example1() {
	var count int

	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}

func example2() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", count)
}

func example3() {
	var onceA, onceB sync.Once
	var initB func()

	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) } // 1. 2가 호출될 때까지 진행되지 않는다.
	onceA.Do(initA)                    // 2. 데드락 발생
}
