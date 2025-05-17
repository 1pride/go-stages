package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"os"
	"sync"

	"exercise/cString"
)

/*
Define an interface for shapes with Area()
Read/write JSON to a file
Print numbers using goroutines
Create a custom string utility package
*/

type Square struct {
	side float64
}

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

type randomN struct {
	I int     `json:"integers"`
	F float64 `json:"floats"`
}

type Shapes interface {
	Area() float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}

func createAndWriteToFile(input string) {
	data, err := os.Create("nums.json")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer data.Close()

	_, err = data.WriteString(input)
	if err != nil {
		log.Fatal("Error writing file", err)
	}
}

func main() {

	circle := Circle{3.5}
	rect := Rectangle{3, 4}
	square := Square{4}

	//fmt.Println(circle.Area())
	//fmt.Println(rect.Area())
	//fmt.Println(square.Area())

	i1 := Shapes(circle)
	i2 := Shapes(rect)
	i3 := Shapes(square)
	fmt.Println("circle:", i1.Area())
	fmt.Println("rectangle:", i2.Area())
	fmt.Println("square:", i3.Area())

	var multiNums []randomN
	for r := range 10 {
		r++
		multiNums = append(multiNums, randomN{I: rand.IntN(200), F: rand.Float64()})
	}

	input, _ := json.MarshalIndent(multiNums, "", "")
	createAndWriteToFile(string(input))

	var wg sync.WaitGroup
	nums := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			nums <- rand.IntN(200)
			wg.Done()
		}()
		b := fmt.Sprintf("goroutines nums: %d", <-nums)
		cString.Blue(b)
	}

	_, err := os.ReadFile("nums.json")
	cString.CheckErr(err)

	wg.Wait()
	//close(nums)

	cString.CountVowels("hello world")

	cString.CapitalizeAndRemoveSpaces("hello\t\tworLd now")
}
