package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
Loop:
	for i := 1; i < 100; i++ {
		for j := 0; j <= 10; j++ {
			if i*j > 3 {
				break Loop
			}
		}
	}

	arr := [3]int{1, 2, 3}
	for k, v := range arr {
		fmt.Println(k, v)
	}
}
