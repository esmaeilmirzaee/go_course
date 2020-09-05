package main

import (
	"fmt"
)

const (
	First       = iota // First is going to be exposed
	secondScore        // naming convention for const type is using camel case. If the const should be exported to outside the first letter must be capitalised.

)

func main() {
	a := 10
	b := 3
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)  // exclusive or
	fmt.Println(a &^ b) // in not
	// Bit shifting 3 times to left and right
	a = 8
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0
	// There are two types of complex number
	// complex64 and complex128
	var n complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", n, n)
	m := 0 + 5.2i
	fmt.Println(n + m)
	fmt.Println(n - m)
	fmt.Println(n * m)
	fmt.Println(n / m)
	fmt.Printf("%v, %T\n", real(n), real(n))
	fmt.Printf("%v, %T\n", imag(n), imag(n))
	m = complex(5, 12)
	fmt.Printf("%v, %T\n", m, m)
	s := "This is string"
	by := []byte(s)
	fmt.Printf("%v, %T\n", by, by)
	r := 'a' // rune
	fmt.Printf("%v, %T\n", r, r)
	fmt.Println(First, secondScore)

	// Arrays
	arr :=
		[3]int{96, 97, 90}
	fmt.Println(arr)
	grades := [...]int{89, 85, 86} // big enough to hold my data
	fmt.Println(grades[0])
	grades[0] = 98
	fmt.Println(grades[0])
	// By default arrays clone in Golang
	anotherGrades := grades
	anotherGrades[1] = 93
	fmt.Println(grades)
	fmt.Println(anotherGrades)
	// Pointer copy
	pointerGradesCopy := &grades
	pointerGradesCopy[0] = 90
	fmt.Println(grades)
	fmt.Println(*pointerGradesCopy)
	// some useful methods
	fmt.Println(len(grades))
	fmt.Println(cap(grades))
	anArray := make([]string, 10, 100) // create slice with capacity 10 and length 100
	// slice
	fmt.Println(grades[:])
	fmt.Println(grades[0:1])
	fmt.Println(grades[1:])
	fmt.Println(grades[:2])
	fmt.Println(grades[0:2])
	someGrades := []int{1}
	someGrades = append(someGrades, 2, 3, 4, 5)
	someGrades = append(someGrades, []int{1, 2, 3}...)
	fmt.Println(someGrades)
	fmt.Println(cap(someGrades))
	// stack operations
	// shift
	var z = []int{1, 2, 3, 4, 5}
	var x = z[1:]
	fmt.Println(x)
	x = z[:len(z)-1]
	fmt.Println(x)
	x = append(z[2:], z[3:]...)
	fmt.Println(x)

}
