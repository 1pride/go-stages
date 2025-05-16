package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name  string
	Age   int
	Alive bool
}

func main() {
	allInfo := User{
		Name:  "John",
		Age:   20,
		Alive: false,
	}
	fmt.Println(allInfo)

	fmt.Println(nameAndAge(allInfo.Name, allInfo.Age))

	fmt.Println("sum of numbers: ", sumOfNumbers(5, 15))

	slice := []string{"one", "two", "three", "three", "three", "two", "three", "two", "two", "TWO", "THREE", "tWo", "THREE", "two"}
	word, count := countWords(slice)
	fmt.Printf("word %s was found %d times\n", word, count)
}

func nameAndAge(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func sumOfNumbers(x, y int) int {
	sum := x + y
	return sum
}

func countWords(phrase []string) ([]string, int) {
	// final result variables
	var finalCount int
	var mostFrequentWords []string

	// map to store inputs
	storing := make(map[string]int)

	// range over slice to get all elements
	for _, word := range phrase {
		// lowering all cases to avoid mixed chars
		lower := strings.ToLower(word)

		// updating map count
		storing[lower]++

		// if _, ok := m[lower]; ok {
		// 	m[lower]++
		// } else {
		// 	m[lower] = 1
		// }

		// final check for getting the values and update to return the values at the end
		if storing[lower] > finalCount {
			finalCount = storing[lower]
			mostFrequentWords = []string{lower}
		} else if storing[lower] == finalCount { // in case has a tier of words
			mostFrequentWords = append(mostFrequentWords, lower)
		}

	}

	return mostFrequentWords, finalCount
}
