package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

var mutex sync.Mutex

func main() {
	var counter int64 = 0
	wg.Add(100)

	go func() {
		for i := 0; i < 100; i++ {
			go func() {
				atomic.AddInt64(&counter, 1)
				wg.Done()
			}()
		}
	}()

	wg.Add(1)
	go func() {
		for atomic.LoadInt64(&counter) < 100 {
			runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&counter))
		}

		fmt.Println("AQUIIIIIIIIIIii")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(counter)

	fmt.Println(runtime.NumGoroutine())
}
