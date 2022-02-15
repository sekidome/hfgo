package main

import "fmt"

func main() {
	var myInt int
	var myIntPointer *int // define a variable as a pointer to an int type
	myInt = 42
	myIntPointer = &myInt      // assign the pointer to an actual int type value ( read & as points to)
	fmt.Println(*myIntPointer) // access the value of the pointer via * ( read * as value behind)
}
