package main

import (
	"fmt"
)

//Original code https://play.golang.org/p/j-EA6003P0
func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
