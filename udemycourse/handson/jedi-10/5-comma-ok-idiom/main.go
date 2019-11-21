//Show the comma ok idiom starting with this code (https://play.golang.org/p/qh2ywLB5OG).

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 1
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
