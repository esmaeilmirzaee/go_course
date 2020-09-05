package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello Function")
	// sayGreeting("Hello ", "Esmaeil")
	// sum("The sum is ", 1, 2, 3, 4, 5)
	// res := sumInt("The sum is ", 1, 2, 3, 4, 5)
	// fmt.Println(res)
	// resP := sumP("The sum is ", 1, 2, 3, 4, 5)
	// fmt.Println(resP, *resP)
	// resR := sumR("The sum is ", 1, 2, 3, 4, 5)
	// fmt.Println(resR)
	// // fmt.Println(divide(5.0, 0.0))
	// d, err := divideE(5.0, 0.0)
	// fmt.Println(d, err)

	func() {
		fmt.Println("Anonymous function")
	}()

	// asynchronous function
	for i := 0; i < 5; i++ {
		func() {
			fmt.Println(i)
		}()
	}

	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}
	// functions as types
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0 {
			return 0.0, fmt.Errorf("Divided by zero")
		}
		return a / b, nil
	}
	d, err := divide(5.0, 3.0)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(d)
}

func sayGreeting(msg, name string) {
	fmt.Println(msg, name)
}

func sum(msg string, values ...int) {
	fmt.Println("Calculating")
	sum := 0
	for _, v := range values {
		sum += v
	}
	fmt.Println(msg, sum)
}

func sumInt(msg string, values ...int) int {
	fmt.Println("Returning result")
	var res int
	for _, v := range values {
		res += v
	}
	return res
}

func sumP(msg string, values ...int) *int {
	fmt.Println("Pointer result")
	var res int
	for _, v := range values {
		res += v
	}
	return &res
}

// syntactic suger
func sumR(msg string, values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

//
func divide(a, b float64) float64 {
	return a / b
}

// handling error idiomatic way
func divideE(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("divide by zero")
	}
	return a / b, nil
}
