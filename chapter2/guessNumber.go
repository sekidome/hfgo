package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds) // this "seeds" the random number generator. -> In go, u need to decide on the way u want to create the pseudorandom numbers.
	ranNum := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100. \n Can u guess it?")
	success := false
	for i := 1; i <= 10; i++ {
		fmt.Print("Guess a number:")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess > ranNum {
			fmt.Println("Oops, you guessed to high.", 10-i, "guesses left.")
		} else if guess < ranNum {
			fmt.Println("Oops, you guessed to low", 10-i, "guesses left.")
		} else {
			fmt.Println("You guessed right on the", i, ". guess.")
			success = true
			break
		}
	}
	if !success {
		fmt.Println("Sorry, time is up. The number was", ranNum)
	}
}
