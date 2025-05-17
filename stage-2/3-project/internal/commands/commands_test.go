package commands_test

import (
	"os"
	"reflect"
	"testing"

	"todo-app/internal/commands"
)

// TODO: learn later how to test properly, kinda odd for now
func TestJSONTaskManager_AddTask(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"simple task", "Hello world", "[ ]Hello world"},
		{"another task", "Read Go docs", "[ ]Read Go docs"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temp file to isolate test data
			tmpFile := "test-todo.json"
			defer os.Remove(tmpFile)

			err := os.WriteFile(tmpFile, []byte("{}"), 0644)
			if err != nil {
				t.Fatal("Error creating temp file", err)
				return
			}

			manager := &commands.JSONTaskManager{JSONFilePath: tmpFile}
			manager.AddTask(tt.input)

			if manager.Storage.Todo == nil {
				t.Fatal("expected to be initialized, got nil")
			}

			got := manager.ListTasks()[1]

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}

func TestJSONTaskManager_ListTasks(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		want        map[int]string
	}{
		{
			name:        "List empty task map",
			fileContent: `{"Todo":{}}`,
			want:        map[int]string{},
		},
		{
			name:        "List single task",
			fileContent: `{"Todo":{"1":"Buy milk"}}`,
			want:        map[int]string{1: "Buy milk"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temp file to isolate test data
			tmpFile := "test-todo.json"
			defer os.Remove(tmpFile)

			err := os.WriteFile(tmpFile, []byte(tt.fileContent), 0644)
			if err != nil {
				t.Fatalf("failed to write test file: %v", err)
			}

			manager := &commands.JSONTaskManager{JSONFilePath: tmpFile}
			got := manager.ListTasks()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONTaskManager_DeleteTask(t *testing.T) {
	tests := []struct {
		name       string
		input      int
		expectGone bool
	}{
		{"delete existing task", 1, true},
		{"fail delete non-existent task", 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temp file to isolate test data
			tmpFile := "test-todo.json"
			defer os.Remove(tmpFile)

			err := os.WriteFile(tmpFile, []byte(`{"Todo":{"1":"Buy milk"}}`), 0644)
			if err != nil {
				t.Fatalf("Error writing test file: %v", err)
			}

			manager := &commands.JSONTaskManager{JSONFilePath: tmpFile}
			manager.DeleteTask(tt.input)

			gotMap := manager.ListTasks()
			_, exists := gotMap[tt.input]

			if tt.expectGone && exists {
				t.Errorf("expected task %d to be deleted, but it still exists", tt.input)
			}
			if !tt.expectGone && exists {
				t.Errorf("expected task %d to still exist, but it was deleted", tt.input)
			}
		})
	}
}

func TestShowCommands(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"show commands", "add -task=\"TASK 1\"\ndel -id=1\nlist\ndone -id=1\n-help"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commands.ShowCommands()
		})
	}
}

func TestJSONTaskManager_ToggleCheck(t *testing.T) {
	tests := []struct {
		name string
		id   int
		want string
	}{
		{"toggle check", 1, "[✅]Buy milk"},
		{"toggle uncheck", 2, "[ ]Buy milk"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile := "test-todo.json"
			defer os.Remove(tmpFile)

			err := os.WriteFile(tmpFile, []byte(`{"Todo":{"1":"[ ]Buy milk","2":"[✅]Buy milk"}}`), 0644)
			if err != nil {
				t.Fatal("Error creating temp file", err)
				return
			}

			manager := &commands.JSONTaskManager{JSONFilePath: tmpFile}
			manager.ToggleCheck(tt.id)

			got := manager.ListTasks()[1]
			got2 := manager.ListTasks()[2]

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %s, want %s", got, tt.want)
			}
			if !reflect.DeepEqual(got2, tt.want) {
				t.Errorf("got2 %s, want %s", got2, tt.want)
			}

		})
	}

}
