// if package main is declear here functions, arrays, and variable can be accessed in the main file for me
// is steve.go
// and any global variable or array in steve.go can be accessed here
package main

import (
	"fmt"
)

func SayHello(hello string) {
	fmt.Printf("Hello world from %v \n", hello)
}

var fog = []string{"fred", "ben", "chuks", "mark", "mathew"}

type bill struct {
	// this is the structure of struct that can take map which as object, array and string
	name string
	item map[string]float32
	tag  float32
}

func createBill(name string) bill {
	// this function will be used to create new struct that are array with object
	b := bill{
		name: name,
		item: map[string]float32{"rice": 30.00, "bean": 20.3, "corn": 45.6},
		tag:  0,
	}
	return b
}

// recevicer functions
// this are used to mapulate the struct it is just like laravel resource or collection
func (b bill) format() string {
	// (b bill) b  is referencing bill struct

	fs := "list of good \n"
	total := float32(0)

	for index, value := range b.item {

		fs += fmt.Sprintf("%v and price:%0.1f \n", index, value)
		total += value
	}

	fs += fmt.Sprintf("total %-7v :%0.1f  \n", " ", total)

	return fs
}
