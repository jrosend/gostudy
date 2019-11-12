package main

import "fmt"

func main() {
	channel := make(chan int)

	go send(channel)
	receive(channel)
	fmt.Println("Ready to exit")
}

func send(channel chan<- int) {
	channel <- 42
}

func receive(channel <-chan int) {
	fmt.Println(<-channel)
}
