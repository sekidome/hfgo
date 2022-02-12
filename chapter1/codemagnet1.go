package main

import (
	"fmt"
)

func main() {
	var originalCount int = 10 /*long var decleration*/
	eatenCount := 4            /*shortform -> doesn't work outside of function*/

	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")
}
