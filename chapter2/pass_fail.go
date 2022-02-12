package main

import (
	"fmt"
)

// this is a comment
func main() {
	testResult := 56
	if testResult > 60 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
