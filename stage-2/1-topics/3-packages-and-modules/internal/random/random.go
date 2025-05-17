package random

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"math/rand"
)

func Guess() {
	number := rand.Intn(100)
	fmt.Println(number)
	var num int

	for {
		fmt.Print("Guess a number: ")
		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			log.Fatal(err)
		}

		if num == number {
			color.Green("You guessed it!")
			break
		}
		if num > number {
			color.Red("Too high!")
		}
		if num < number {
			color.Red("Too low!")
		}
	}
}
