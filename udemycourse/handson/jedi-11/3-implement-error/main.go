package main

import "log"

type customErr struct {
}

func (err *customErr) Error() string {
	return "An error occurred"
}

func main() {
	err := &customErr{}

	foo(err)
}

func foo(err error) {
	log.Println(err)
}
