package main

import (
	"fmt"
	"math/rand"
        "unicode"
)

func main() {

	fmt.Println("Welcome! Guess the letter from 'a' to 'z' that I'm thinking of!")

	min := 97 // 'a' in ASCII
	max := 122 // 'z' in ASCII
        pick := int32(rand.Intn(max-min+1) + min)

runner:
	for {
		fmt.Print("Guess: ")
		var input string
		fmt.Scanln(&input)

                runes := []rune(input)

                if len(runes) == 0 {
                        fmt.Println("Input Error! Enter a real character.")
                        continue
                }

                if !unicode.IsLetter(runes[0]) {
                        fmt.Println("Input Error! Please enter an English letter.")
                        continue
                }
                
                guess := unicode.ToLower(runes[0])

		switch {
		case guess == pick:
			fmt.Println("Correct! You WIN!")
			break runner
		case guess < pick:
			fmt.Println("Too low :(")
		case guess > pick:
			fmt.Println("Too high :O")
		}
	}
}
