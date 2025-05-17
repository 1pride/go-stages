package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"strings"
)

type taskInfo struct {
	Todo map[int]string
}

type JSONTaskManager struct {
	JSONFilePath string
	Storage      taskInfo
}

// findAvailable returns the smallest positive integer that is not a key in the provided map.
func findAvailable(m map[int]string) int {
	for i := 1; ; i++ {
		if _, exists := m[i]; !exists {
			return i
		}
	}
}

// checkErr can be used to simplify all if err != nil {log.Fatal}
func checkErr(err error, msg ...string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

// load reads tasks from the JSON file at JSONFilePath and initializes the Storage field with the data.
// if the file does not exist, it initializes Storage with an empty taskInfo structure.
func (j *JSONTaskManager) load() {
	data, err := os.ReadFile(j.JSONFilePath)
	if errors.Is(err, os.ErrNotExist) {
		j.Storage = taskInfo{Todo: make(map[int]string)}
		return
	}
	checkErr(json.Unmarshal(data, &j.Storage), "unmarshal error")
}

// save writes the current state of Storage to the JSON file specified in JSONFilePath, formatting the output with indentation.
func (j *JSONTaskManager) save() {
	data, err := json.MarshalIndent(j.Storage, "", "  ")
	checkErr(err, "marshal error")
	checkErr(os.WriteFile(j.JSONFilePath, data, 0644), "write error")
}

// AddTask adds a new task to the storage with an auto-generated ID and persists the updated data to the JSON file.
func (j *JSONTaskManager) AddTask(input string) {
	j.load()
	if j.Storage.Todo == nil {
		j.Storage.Todo = make(map[int]string)
	}
	id := findAvailable(j.Storage.Todo)
	j.Storage.Todo[id] = "[ ]" + input
	j.save()
}

// DeleteTask removes a task with the specified index from the storage and saves the updated data to the JSON file.
func (j *JSONTaskManager) DeleteTask(index int) {
	j.load()
	delete(j.Storage.Todo, index)
	j.save()
}

// ToggleCheck toggles the checked state of a task at the specified index in the storage and saves the updated data.
func (j *JSONTaskManager) ToggleCheck(index int) {
	j.load()
	task := j.Storage.Todo[index]
	if strings.Contains(task, "[ ]") {
		j.Storage.Todo[index] = strings.Replace(task, "[ ]", "[✅]", 1)
	} else if strings.Contains(task, "[✅]") {
		j.Storage.Todo[index] = strings.Replace(task, "[✅]", "[ ]", 1)
	}
	j.save()
}

// ListTasks retrieves and prints the list of tasks sorted by their IDs.
func (j *JSONTaskManager) ListTasks() map[int]string {
	j.load()
	for _, key := range slices.Sorted(maps.Keys(j.Storage.Todo)) {
		fmt.Println(key, j.Storage.Todo[key])
	}
	return j.Storage.Todo
}

// ShowCommands displays a list of available CLI commands with their respective syntax.
func ShowCommands() {
	fmt.Println(`----------------- CLI COMMANDS -----------------
| go run ./cmd/main.go add -task="TASK 1"      |
| go run ./cmd/main.go done -id=1              |
| go run ./cmd/main.go del -id=1               |
| go run ./cmd/main.go list                    |
| go run ./cmd/main.go help                    |
----------------- CLI COMMANDS -----------------`)
}
