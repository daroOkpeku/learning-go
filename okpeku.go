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
	tip  float32
}

func createBill(name string) bill {
	// this function will be used to create new struct that are array with object
	b := bill{
		name: name,
		item: map[string]float32{"rice": 30.00, "bean": 20.3, "corn": 45.6},
		tip:  0,
	}
	return b
}

// recevicer functions
// this are used to mapulate the struct it is just like laravel resource or collection
// and i added the the * for defereancing to update the original copy
func (b *bill) format() string {
	// (b bill) b  is referencing bill struct

	fs := "list of good \n"
	total := float32(0)

	for index, value := range b.item {

		fs += fmt.Sprintf("%v and price:%0.1f \n", index, value)
		total += value
	}
	fs += fmt.Sprintf("tip %-7v :%0.1f  \n", " ", b.tip)
	fs += fmt.Sprintf("total %-7v :%0.1f  \n", " ", total+b.tip)

	return fs
}

func (b *bill) AddItems(index string, value float32) {
	// to update items in the struct you have to use pointer for the struct
	// this is how to do it (b *bill) golang automaticatly dereference the struct from the orignal copy
	b.item[index] = value

	// this is allow way of deferencing in go
	// (*b).item[index] = value and b.item[index] = value is allowed in go
}

func (b *bill) updateTip(value float32) {
	b.tip = value
}

func (b *bill) addItem() {

}
