package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(4)
	go func() {
		fmt.Println("goRoutine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("goRoutine 2")
		wg.Done()
	}()

	go func() {
		fmt.Println("goRoutine 3")
		wg.Done()
	}()

	go func() {
		fmt.Println("goRoutine 4")
		wg.Done()
	}()

	wg.Wait()
}
