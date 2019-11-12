package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)
	receive(even, odd, quit)

	fmt.Println("Exiting...")
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v, _ := <-even:
			fmt.Println("Even", v)
		case v, _ := <-odd:
			fmt.Println("Odd", v)
		case _, ok := <-quit:
			fmt.Println("Quit", ok)
			return
		}
	}
}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	quit <- 0
}
