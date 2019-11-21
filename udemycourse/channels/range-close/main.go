package main

import "fmt"

func main() {
	channel := make(chan int)

	go send(channel)

	for v := range channel {
		fmt.Println(v)
	}

	fmt.Println("Ready to exit")
}

func send(channel chan<- int) {

	for i := 0; i < 100; i++ {
		channel <- i
	}
	close(channel)
}

func receive(channel <-chan int) {
	fmt.Println(<-channel)
}
