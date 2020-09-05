package main

import "fmt"

func main() {
	fmt.Println("Method")
	g := greeter{
		name: "Go",
		msg:  "Hello",
	}
	g.greet()
}

type greeter struct {
	name string
	msg  string
}

// Methods
func (g greeter) greet() {
	fmt.Println(g.name, g.msg)
}
