package main

import (
	"flag"
	"fmt"
	"os"
	"todo-app/internal/commands"
)

func check(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		return
	}
}

func main() {
	manager := &commands.JSONTaskManager{JSONFilePath: "data/todo.json"}

	addTask := flag.NewFlagSet("add", flag.ExitOnError)
	task := addTask.String("task", "", "insert a valid name to add a new task")

	toggleCheck := flag.NewFlagSet("done", flag.ExitOnError)
	idCheck := toggleCheck.Int("id", -1, "insert a valid name to add a new task")

	listTask := flag.NewFlagSet("list", flag.ExitOnError)

	delTask := flag.NewFlagSet("del", flag.ExitOnError)
	id := delTask.Int("id", -1, "insert a valid name to add a new task")

	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Invalid command, try go run main.go add, del, list, done or help")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--add", "add", "-A":
		err := addTask.Parse(os.Args[2:])
		check("Error adding a task", err)

		manager.AddTask(*task)
		manager.ListTasks()
	case "--list", "list", "-L":
		err := listTask.Parse(os.Args[2:])
		check("Error listing tasks", err)

		manager.ListTasks()
	case "--delete", "del", "-D":
		err := delTask.Parse(os.Args[2:])
		check("Error deleting a task at parse", err)

		if *id == -1 {
			fmt.Println("Error deleting a task, id not provided")
			os.Exit(1)
		}

		ok := manager.ListTasks()
		_, exist := ok[*id]
		if exist {
			fmt.Printf("Task %v deleted, content: %v\n", *id, ok[*id])
			manager.DeleteTask(*id)
		} else {
			fmt.Println("Error deleting a task, wrong id")
			os.Exit(1)
		}
	case "--done", "done":
		err := toggleCheck.Parse(os.Args[2:])
		check("Error toggling a task", err)

		if *idCheck == -1 {
			fmt.Println("Error deleting a task, id not provided")
			os.Exit(1)
		}

		ok := manager.ListTasks()
		_, exist := ok[*idCheck]
		if exist {
			manager.ToggleCheck(*idCheck)
			fmt.Printf("Task ID: %v new status, content: %v\n", *idCheck, ok[*idCheck])
		} else {
			fmt.Println("Error checked/unchecked, wrong id")
			os.Exit(1)
		}
	case "--help", "help", "-H":
		err := helpCmd.Parse(os.Args[2:])
		check("Error listing a task", err)

		commands.ShowCommands()
	default:
		fmt.Println(`Invalid command, try: help`)
		os.Exit(1)
	}
}
