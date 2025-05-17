package main

import (
	"errors"
	"fmt"
	"log"
)

var errFileNotFound = errors.New("file not found")

// FileError Define a custom error type for file errors
type FileError struct {
	Op  string // Operation that caused the error
	Err error  // Underlying error
}

// Implement the Error() method for FileError
func (e *FileError) Error() string {
	return fmt.Sprintf("file error during %s: %v", e.Op, e.Err)
}

// Unwrap Implement the Unwrap() method to allow unwrapping the underlying error
func (e *FileError) Unwrap() error {
	return e.Err
}

// Simulate a function that reads a file and may return an error
func readFile(filename string) error {
	// Simulating a file not found error
	if filename == "" {
		return &FileError{
			Op:  "read",
			Err: errFileNotFound,
		}
	}
	// Simulating a successful read
	return nil
}

// Simulate a function that processes a file and handles errors
func processFile(filename string) error {
	err := readFile(filename)
	if err != nil {
		// Wrap the error with additional context
		return fmt.Errorf("failed to process file %s: %w", filename, err)
	}
	return nil
}

func main() {
	//Function	Purpose	Uses                    Unwrap()?	    Matching Rule
	//Unwrap	Gets directly wrapped error  	✅ One level	N/A
	//Is	    Checks if error matches target	✅ Recursive	== comparison
	//As	    Checks if error is of a type	✅ Recursive	type assertion

	nonWrapped := errors.New("non wrapped error")
	base := errors.New("base error")

	wrapped := fmt.Errorf("once wrapped error: %w", base)
	fmt.Println(wrapped)
	wrapped2 := fmt.Errorf("twice wrapped error: %w", wrapped)
	fmt.Println(wrapped2)

	//errors.Is walks the error chain by calling Unwrap() recursively until it finds a match using ==.
	if errors.Is(wrapped2, nonWrapped) {
		fmt.Println("contains base error")
	} else {
		fmt.Println("doesn't contain base error")
	}

	//errors.As walks the chain like Is, but uses type assertions instead of value equality.
	myErr := &FileError{"example of using errors.As", base}
	wrappedAs := fmt.Errorf("wrapped error: %w", myErr)

	var target *FileError
	if errors.As(wrappedAs, &target) {
		fmt.Println("Found MyError and target msg: ", target.Op)
	}
	fmt.Println("------------------ example code ------------------")
	// Attempt to process a file with an empty filename

	err := processFile("")
	if err != nil {
		// Check if the error is of the type FileError
		var fileErr *FileError
		if errors.As(err, &fileErr) {
			log.Printf("Caught a FileError: %v", fileErr)
			// You can also access the underlying error
			if underlyingErr := errors.Unwrap(fileErr); underlyingErr != nil {
				log.Printf("Underlying error: %v", underlyingErr)
			}
		} else {
			log.Printf("An unexpected error occurred: %v", err)
		}

		// Check if the error is specifically a "file not found" error
		if errors.Is(err, errFileNotFound) {
			log.Println("Handle file not found error specifically.")
		}
	}

}
