package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Defer, Panic and Recover")
	defer fmt.Println("Start")
	fmt.Println("Middle")
	fmt.Println("End")
	// getRes() continue

	// defer read the data when it defers
	a := "start"
	defer fmt.Println(a)
	a = "end"
	panicSection()
}

// defer
func getRes() {
	res, err := http.Get("https://google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func panicSection() {
	// 1
	// a, b := 1, 0
	// res := a/b
	// fmt.Println(res)

	// 2
	fmt.Println("Start")
	// panic("Panic") due to continue
	fmt.Println("End")

	// 3
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	// err := http.ListenAndServe(":8080", nil) continue
	// if err != nil {
	// panic(err.Error)
	// }

	// 4
	fmt.Println("Start Panicker")
	panicker()
	fmt.Println("End Panicker")
}

func panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error ", err)
		}
	}()
	panic("Something bad happened")
	fmt.Println("End of panic")
}
