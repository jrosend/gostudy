package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	counter := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			x := counter

			runtime.Gosched()

			x++

			counter = x
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
