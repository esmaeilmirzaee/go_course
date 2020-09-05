package main

import (
	"fmt"
)

func main() {
	fmt.Println("If/Else statements")
	if true {
		fmt.Println("Curly brackets are a must...")
	}

	statePopulations := map[string]int{
		"California": 39250012,
	}
	if pop, ok := statePopulations["California"]; ok {
		fmt.Println(pop)
	}
	pop := statePopulations["California"]
	switch {
	case pop < 20000000:
		fmt.Println("Small size city")
	case pop > 20000000 && pop < 25000000:
		fmt.Println("Medium size city")
	}

	var i interface{} = "string"
	switch i.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
