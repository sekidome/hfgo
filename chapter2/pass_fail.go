package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// this is a comment
func main() {
	fmt.Print("Enter a grade:")
	reader := bufio.NewReader(os.Stdin) // reading loop
	//fmt.Println(reader) --> This will keep on printing 0s
	input, err := reader.ReadString('\n') // ReadString method requires an argument with a rune that marks the end of the input
	var status string
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	fmt.Println(grade)
	if grade > 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("The grade", grade, "is", status)
}
