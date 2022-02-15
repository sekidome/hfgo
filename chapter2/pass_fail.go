package main

import (
	"fmt"
	keyboard "hfgo/chapter4"
	"log"
)

// this is a comment
func main() {
	fmt.Print("Enter a grade:")
	var status string
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("The grade", grade, "is", status)
}
