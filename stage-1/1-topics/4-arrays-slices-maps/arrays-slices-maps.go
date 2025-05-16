package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("------------ ARRAYS ------------")
	// arrays have a limited size and for accessing them is good to know they start at 0 (not 1)
	array1 := [3]int{1, 2, 3}
	fmt.Println("accessing specific element", array1[0])

	// is possible to create arrays like that [...]T and Go will implicit the size
	array2 := [...]string{"one", "two", "three"}
	fmt.Println("dynamic size", array2[2])
	fmt.Printf("size of dynamic is %d and len %d\n", cap(array2), len(array2))
	//in case of accessing an index inexistent will get an error index out of bounds...
	// fmt.Println(array2[3])

	// also possibly create an array with pre-defined size and no values (they will start with specific T zero values)
	array3 := [3]bool{}
	fmt.Println("zero value array3", array3)

	//changing array values
	array3[0] = true
	fmt.Println("modified array3", array3)

	// slices works similar to array, but they don't have a specific size, so they grow exponentially as needed
	var slice []float64
	fmt.Println("empty slice", slice)
	fmt.Println("cap", cap(slice))
	fmt.Println("len", len(slice))

	// is possible to insert values to an existent slice
	slice = append(slice, 3.14, 20.1)
	fmt.Println("slice after appending values", slice)
	fmt.Println("cap", cap(slice))
	fmt.Println("len", len(slice))

	fmt.Println("------------ MAP's ------------")
	// IMPORTANT maps aren't ordered in for loops, so be careful

	// map at their default value
	map1 := make(map[int]string)
	fmt.Println(map1)

	// adding single value to map
	map1[4] = "four"
	fmt.Println(map1)

	// adding multiple values to map
	map1 = map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println(map1)

	// accessing map values
	fmt.Println("retrieving map value where the key is 2 ->", map1[2])

	// ranging over a map and printing the key and value
	for k, v := range map1 {
		fmt.Printf("key: %d , value: %s\n", k, v)
	}

	// checking if a value exists in a map using the "comma ok idiom" feature
	if value, ok := map1[3]; ok {
		fmt.Printf("value %v exist? %v in map1\n", value, ok)
	}

	// deleting a specific value from a map
	delete(map1, 1)
	fmt.Println("map after deleting value/key 1 -> ", map1)

	// clear all values from a map
	clear(map1)
	fmt.Println("clear the whole map ->", map1)

	// comparing maps
	map2 := map[int]string{1: "one", 2: "two", 3: "three"}
	if maps.Equal(map1, map2) {
		fmt.Println("map1 and map2 are equal")
	} else {
		fmt.Println("map1 and map2 aren't equal")
	}
}
