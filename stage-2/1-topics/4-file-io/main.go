package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"sync"
)

// if the struct has private fields, they will not be written to file
type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// saveToFile create a file and then write the data provided by the user
func saveToFile(input string) {
	data, err := os.Create("users.json")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer data.Close()

	_, err = data.WriteString(input)
	if err != nil {
		log.Fatal("Error writing file", err)
	}
}

// readFile will read the entire file content and print it to the terminal
func readFile(file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	fmt.Println(string(content))
}

// multiLinesToFile creates a file named "lines.txt" and writes multiple lines of text to it, then prints the total bytes written.
func multiLinesToFile() {
	newFile, err := os.Create("lines.txt")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer newFile.Close()

	lines := []string{"line1", "line2", "line3", "line4", "line5", "line6", "line7", "line8", "line9", "line10"}
	var totalBytes int
	for _, line := range lines {
		// method 1
		//_, err := newFile.WriteString(line + "\n")
		//if err != nil {
		//	log.Fatal("Error writing file", err)
		//}

		//method 2
		bytes, _ := fmt.Fprintln(newFile, line)
		totalBytes += bytes
	}
	fmt.Printf("%d bytes\n", totalBytes)

}

// ReadFileLineByLine reads a file line by line and prints each line to the terminal.
func ReadFileLineByLine(file string) {
	data, err := os.Open(file)
	//if err != nil {
	//	log.Fatal("Error opening file", err)
	//}
	check(err)
	defer data.Close()

	lineScan := bufio.NewScanner(data)
	for lineScan.Scan() {
		fmt.Println(lineScan.Text())
	}

	if scanErr := lineScan.Err(); scanErr != nil {
		fmt.Printf("Error during scanning line by line: %v\n", scanErr)
	}
}

func check(e error, msg ...string) {
	if e != nil {
		log.Fatal(msg, ": ", e)
	}
}

func main() {
	multiUser := []user{
		{Name: "John", Age: 22},
		{Name: "Jane", Age: 23},
		{Name: "Jake", Age: 24},
	}

	//input := user{Name: "John", Age: 22}
	//bytes, _ := json.Marshal(input)
	bytes, _ := json.MarshalIndent(multiUser, "", "")

	saveToFile(string(bytes))

	readFile("users.json")

	multiLinesToFile()

	ReadFileLineByLine("lines.txt")

	// go routines making it concurrently
	data := make(chan int)
	done := make(chan bool)

	var wg sync.WaitGroup

	logFile, err := os.Create("logs.txt")
	check(err, "error at creating file")
	defer logFile.Close()

	// mock producers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			data <- rand.IntN(200)
			wg.Done()
		}()
	}

	// conductor
	go func() {
		wg.Wait()
		done <- true
	}()

	for {
		select {
		case value := <-data:
			_, writeErr := fmt.Fprintln(logFile, value)
			if writeErr != nil {
				done <- false
			}
		case signal := <-done:
			// if the signal is true
			if signal {
				fmt.Println("ALL DONE NICELY")
				return // not returning will lead to "all goroutines are asleep deadlock"
			} else {
				fmt.Println("NOT OKAY")
				return
			}
		}
	}
}
