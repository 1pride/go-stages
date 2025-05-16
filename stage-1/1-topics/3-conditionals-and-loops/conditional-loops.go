package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println("------------ IF's ------------")
	showCaseIfStatement()
	fmt.Println("------------ SWITCH ------------")
	switchStatements("Friday")
	fmt.Println("------------ FOR LOOPS ------------")
	forLoops()
}

// showCaseIfStatement show examples of how if statements work and how useful they can be in different scenarios
func showCaseIfStatement() {
	// a simple if
	if 10 > 1 {
		fmt.Println("10 is bigger than 1")
	}

	// if and else if
	if "equal" != "equal" {
		fmt.Println("not equal")
	} else if "equal" == "equal" {
		fmt.Println("they are equal")
	}

	// if, else if and else
	if true != true {
		fmt.Println("true isn't equal true")
	} else if true == false {
		fmt.Println("true is equal to false")
	} else {
		fmt.Println("true is only equal to true")
	}

	// handling errors
	if err := doSomething(); err != nil {
		log.Print(err) //prints date, time and error msg
	}
}

// doSomething fake an error to show how to handling an error using if err != nil...
func doSomething() error {
	return errors.New("fake error")
}

// switchStatements show examples of how a switch statement works and how useful they can be in different scenarios
func switchStatements(day string) {
	// basic switch
	switch day {
	case "Monday":
		fmt.Println("Today is", day)
	case "Sunday":
		fmt.Println("Today is", day)
	default:
		fmt.Println("Another day")
	}

	// multiple values in one case
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}

	// switch without an expression (acts like chained if?)
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade A")
	case score >= 80:
		fmt.Println("Grade B")
	case score >= 70:
		fmt.Println("Grade C")
	default:
		fmt.Println("Fail")
	}

	// type switch (dynamic typing on interfaces)
	var i any = "hello"
	// var i interface{} = "hello"

	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("unknown type")
	}

}

// forLoops show examples of how for loops works and how useful they can be in different scenarios
func forLoops() {
	// basic loop
	for i := 0; i < 3; i++ {
		fmt.Println("i:", i)
	}

	// while loop
	x := 0
	for x < 3 {
		fmt.Println("x:", x)
		x++
	}

	// infinite loop
	for {
		fmt.Println("unless has a break will continue")
		break // use break or return to stop running for ever
	}

	// for range loop using string as demo (aka runes/chars)
	str := "hello"
	for i, r := range str {
		fmt.Printf("Byte %d: Rune %c\n", i, r)
	}

	// ps: maps are unordered
	// for range with maps
	m := map[int]string{1: "A", 2: "B", 3: "C"}
	for k, v := range m {
		fmt.Printf("Key: %d, Value: %s\n", k, v)
	}

	// for range with channels (useful in goroutines)
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}
}
