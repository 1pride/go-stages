package cString

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

func Red(s string) string {
	return color.RedString(s)
}

func Blue(s string) string {
	fmt.Println(color.BlueString(s))
	return color.BlueString(s)
}

func CheckErr(err error) {
	if err != nil {
		//log.Fatal(Red("ERROR at " + err.Error()))
		panic(Red("ERROR at " + err.Error()))
	}
}

func CountVowels(phrase string) int {
	count := make(map[string]int)
	totalVowels := 0
	lowered := strings.ToLower(phrase)
	for _, r := range lowered {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			count[string(r)]++
		}
	}

	for _, v := range count {
		totalVowels += v
	}

	fmt.Printf("Total Vowels: %d, Vowels A: %d, E :%d, I :%d, O :%d, U :%d\n", totalVowels, count["a"],
		count["e"], count["i"], count["o"], count["u"])
	return totalVowels
}

func CapitalizeAndRemoveSpaces(words string) string {
	words = strings.Replace(words, "\t", " ", -1) // replace tabs per space
	words = strings.Replace(words, "\n", " ", -1) // replace new lines per space
	words = strings.ToLower(words)                //lowering to avoid mixed cases

	var count int    // check how many spaces extras have been found to remove
	var upper string //just to store the uppercased letter and make the code cleaner

	for i, letter := range words {

		if i == 0 { //uppercase first letter
			upper = strings.ToUpper(string(letter))
			words = upper[i:] + words[i+1:]
		}

		if letter == 32 { //32 is the rune code for space
			upper = strings.ToUpper(string(words[i+1])) //after found space, uppercase the next letter
			words = words[:i+1] + upper + words[i+2:]   //concat the previous slice with the uppercased letter and
			// removed the duplicate letter, keeping the remaining letters
		}

		if letter == 32 {
			if words[i+1] == 32 { // if the next element is a space update count
				count++
			}
		}
	}

	words = strings.Replace(words, " ", "", count)
	return words
}
