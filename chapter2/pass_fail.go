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
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // ReadString method requires an argument with a rune that marks the end of the input
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	fmt.Println(grade)
	if grade > 60 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
