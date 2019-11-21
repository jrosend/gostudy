package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fantOutInt(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("Exit...")

}

func populate(c chan<- int) {
	for i := 1; i <= 100; i++ {
		c <- i
	}
	close(c)
}

func fantOutInt(c1 <-chan int, c2 chan<- int) {
	var wg sync.WaitGroup

	for v := range c1 {
		wg.Add(1)
		go func(value int) {
			c2 <- timeConsumingTask(value)
			wg.Done()
		}(v)
	}

	wg.Wait()
	close(c2)
}

func timeConsumingTask(v int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return v * rand.Intn(1000)
}
