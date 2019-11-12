package main

import "fmt"

type person struct {
	speech string
}

func (p *person) speak() {
	fmt.Println(p.speech)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		speech: "Hello, World",
	}

	saySomething(&p)
}
