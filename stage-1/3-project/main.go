package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"slices"
)

// Contact is a data to provide in the contact book
type Contact struct {
	Name        string
	PhoneNumber string
}

// listOfContacts is a slice that works like a save file for now
var listOfContacts []Contact

func main() {
	// flags available
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)

	// add inputs
	name := addCmd.String("name", "", "insert a valid name")
	number := addCmd.String("number", "", "insert a valid number")

	// delete inputs
	delName := delCmd.String("name", "", "insert a valid name")
	delNumber := delCmd.String("number", "", "insert a valid number")

	// pre-defined users to showcase only and debug proposes
	addContact("John", "1")
	addContact("Jane", "2")
	addContact("Jake", "3")

	if len(os.Args) < 2 {
		fmt.Println("Invalid command, try -add, -delete, -list, -help")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--add", "add", "-A":
		err := addCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Error adding a contact", err)
			return
		}

		addContact(*name, *number)
		showAgenda()
	case "--list", "list", "-L":
		err := listCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Error listing a contact", err)
			return
		}

		showAgenda()
	case "--delete", "del", "-D":
		err := delCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Error deleting a contact", err)
			return
		}

		contact, err := deleteContact(*delName, *delNumber)
		if err != nil {
			fmt.Println("Error deleting a contact in function", err)
			return
		}

		fmt.Println("Contact:", contact)
		showAgenda()
	case "--help", "help", "-H":
		err := helpCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Error listing a contact", err)
			return
		}

		showCommands()
	default:
		fmt.Println("Invalid command, try -add, -delete, -list, -help")
		os.Exit(1)
	}
}

// addContact is a simple way to add contact to a []Contact
func addContact(name, number string) []Contact {
	listOfContacts = append(listOfContacts, Contact{name, number})
	return listOfContacts
}

// showAgenda return a list of ALL contacts
func showAgenda() []Contact {
	fmt.Println(listOfContacts)
	return listOfContacts
}

// deleteContact should remove a contact from the list []Contact in case of both name and number matches
func deleteContact(name, number string) (string, error) {
	for i, c := range listOfContacts {
		if (name == c.Name && number != c.PhoneNumber) || (name != c.Name && number == c.PhoneNumber) {
			return fmt.Sprintf("Name: %s or Number: %s aren't correct!", name, number), errors.New("invalid name or number")
		}

		if name == c.Name && number == c.PhoneNumber {
			r := slices.Delete(listOfContacts, i, i+1)
			listOfContacts = listOfContacts[:0]
			listOfContacts = append(listOfContacts, r...)
		}
	}

	return fmt.Sprintf("%s Number: %s Deleted Successful", name, number), nil
}

// showCommands will show ALL available commands and their usage.
func showCommands() {
	fmt.Println("List of available commands and their usage (add and delete are Case Sensitive)")
	fmt.Println("go run main.go del -name=John -number=1")
	fmt.Println("go run main.go add -name=James -number=1")
	fmt.Println("go run main.go list")
}
