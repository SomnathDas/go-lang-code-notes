package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func simpleGuessingGame() {
	const answer = 65
	reader := bufio.NewReader(os.Stdin)

Loop:
	for {
		fmt.Println("Guess the answer: ")
		guessed, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		guessedInt, errGuessed := strconv.Atoi(strings.Split(guessed, "\n")[0])

		if errGuessed != nil {
			fmt.Println("Please Enter Numeric Integer Values Only!")
			continue
		}

		fmt.Printf("You gussed: %v\n", guessedInt)

		switch {
		case guessedInt < answer:
			fmt.Println("Colder!")
		case guessedInt > answer:
			fmt.Println("Hotter!")
		case guessedInt == answer:
			fmt.Println("You Won!")
			break Loop
		}

	}
}
