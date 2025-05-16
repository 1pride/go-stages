package main

import "fmt"

// initializing a struct with name and age fields
type user struct {
	name string
	age  int
}

type secretPassword struct {
	user
	password string
}

func main() {
	fmt.Println("------------ Normal Struct ------------")
	// giving values to the user struct with implicit field names (they need all be written otherwise you get an error)
	person1 := user{"John", 20}
	fmt.Println(person1)

	// is possible to omit a field when you aren't using the example above
	person2 := user{
		// Name: "Jane",
		age: 19,
	}
	// accessing values with a specific field
	fmt.Println(person2.age)

	fmt.Println("------------ Anonymous Struct ------------")
	//anonymous struct
	anonStruct := struct {
		text   string
		isAnon bool
	}{
		"Anonymous struct",
		true,
	}
	fmt.Println(anonStruct)

	fmt.Println("------------ Function returning struct ------------")
	fmt.Println(newPerson("Jake", 18))

	fmt.Println("------------ Struct Embedding ------------")
	login := secretPassword{
		user: user{
			name: "Mark",
			age:  30,
		},
		password: "hashedPWD",
	}
	fmt.Println(login)
	// accessing embed fields and checking if they are the same
	fmt.Println(login.user.name == login.name) //true
}

func newPerson(name string, age int) user {
	newUser := user{name, age}
	return newUser
}
