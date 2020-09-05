package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	Canfly   bool
}

type User struct {
	email string `required max:"100"`
	name  string
}

func main() {
	fmt.Println("Map")
	statePopulations := map[string]int{
		"California": 392500017,
		"New York":   19745289,
	}
	// The order of elements in map is not guaranteed.
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["California"])
	fmt.Println(statePopulations["Georgia"])
	statePopulations["Georgia"] = 10123120
	fmt.Println(statePopulations["Georgia"])
	delete(statePopulations, "Georgia")
	popPulation, ok := statePopulations["Ohio"]
	popGeorgiaPulation, okGeorgia := statePopulations["Georgia"]
	popCaliforniaPulation, okCalifornia := statePopulations["California"]
	fmt.Println(popPulation, ok)
	fmt.Println(popGeorgiaPulation, okGeorgia)
	fmt.Println(popCaliforniaPulation, okCalifornia)

	fmt.Println("Struct")
	aDoctor := Doctor{number: 3, actorName: "Esmaeil MIRZAEE", companions: []string{
		"Esmaeil MIRZAEE",
		"Esmaeil MIRZAEE",
	}}
	fmt.Printf("%v,\n %T\n", aDoctor.companions, aDoctor.companions)
	tomDoctor := struct{ name string }{name: "Esmaeil MIRZAEE"}
	esmDoctor := tomDoctor
	esmDoctor.name = "Tom Johnson"
	fmt.Println(tomDoctor)
	fmt.Println(esmDoctor)

	// Go doesn't have inheritance (isـa) instend it's composition (has_a)
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48.1
	b.Canfly = false
	fmt.Println(b)

	t := reflect.TypeOf(User{})
	field, _ := t.FieldByName("email")
	fmt.Println(field.Tag)
}
