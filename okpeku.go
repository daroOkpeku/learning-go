// if package main is declear here functions, arrays, and variable can be accessed in the main file for me
// is steve.go
// and any global variable or array in steve.go can be accessed here
package main

import "fmt"

func SayHello(hello string) {
	fmt.Printf("Hello world from %v \n", hello)
}

var fog = []string{"fred", "ben", "chuks", "mark", "mathew"}
