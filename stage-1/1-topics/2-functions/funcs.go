package main

import (
	"errors"
	"fmt"
)

// Default main go func where run all programs
func main() {
	fmt.Println("------------ Normal Func ------------")
	Regular()
	fmt.Println("------------ Func with Params ------------")
	params("john", 20)

	fmt.Println("------------ Func with return ------------")
	// func with return's to show on console they need to be inside a print?!
	fmt.Println(onlyReturn())

	fmt.Println("------------ Func with params and return ------------")
	fmt.Println(paramsAndReturn("Jane"))

	fmt.Println("------------ Func with multi return's ------------")
	fmt.Println(multiReturn())

	fmt.Println("------------ Func with multi params and return's ------------")
	fmt.Println(multiParamsAndReturn("jake", 15))

	fmt.Println("------------ Func with variadic param ------------")
	variadic(1, 2, 3, 4)
}

// Comments above functions work as documentation
// Function, Variable, Struct... which start with Uppercase = Global

// Regular just prints a message
func Regular() {
	fmt.Println("Function with no param's and return's")
}

// params will print a message with the name and age chosen by the "user"
func params(name string, age int) {
	fmt.Printf("Hello my name is :%s and i've %d years old\n", name, age)
}

// onlyReturn will just return an error which they are really useful for custom errors
func onlyReturn() error {
	return errors.New("returning a custom error")
}

// multiReturn are useful to return a safety Type and error (if has one) otherwise nil means no errors
func multiReturn() (int, error) {
	return 10, nil
}

// paramsAndReturn need a provided param and should return the explicit Type
func paramsAndReturn(name string) string {
	return "demonstration name" + name
}

// multiParamsAndReturn need to be provided all params and also return both of types
func multiParamsAndReturn(name string, age int) (string, error) {
	return fmt.Sprintf("No errors found, so my name is %s and i have %d years old", name, age), nil
}

// variadic accept 0 or more args e.g. variadic() or variadic(1, 2, 3, 4) and create a slice of []T
func variadic(nums ...int) {
	fmt.Printf("nums provided %v\n", nums)
}
