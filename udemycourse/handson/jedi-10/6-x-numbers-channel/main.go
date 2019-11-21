//write a program that
//puts 100 numbers to a channel
//pull the numbers off the channel and print them

package main

import "fmt"

func main() {
	c := getChannel(50)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Bye")

}

func getChannel(count int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i <= count; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
