package main

import "fmt"

func first() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cs ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
