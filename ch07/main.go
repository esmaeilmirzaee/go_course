package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	// first()
	// second()
	// third()
	fourth()
}

func first() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(&a, b)
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, b)
	fmt.Println(&a, b)
	fmt.Println(a, *b)
	*b = 10
	fmt.Println(a, b)
	fmt.Println(&a, b)
	fmt.Println(a, *b)
}

func second() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v %p %p\n", a, b, c)
}

type myStruct struct {
	foo int
}

func third() {
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)
	ms = new(myStruct)
	fmt.Println(ms)
	(*ms).foo = 30
	fmt.Println((*ms).foo)
	fmt.Println("A help from compiler--syntactic suger")
	ms.foo = 10
	fmt.Println(ms.foo)
}

func fourth() {
	a := [3]int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)
	// slice
	c := []int{1, 2, 3}
	d := c
	fmt.Println(c, d)
	c[1] = 42
	fmt.Println(c, d)
	// Map
	e := map[string]string{"foo": "bar", "baz": "buz"}
	f := e
	fmt.Println(e, f)
	e["foo"] = "qux"
	fmt.Println(e, f)
}
