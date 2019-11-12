package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var mutex sync.Mutex

func main() {
	counter := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()

			x := counter
			x++
			counter = x
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
