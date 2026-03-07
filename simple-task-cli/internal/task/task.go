package task

import (
	"fmt"
	"encoding/json"
	"os"
)

type TaskStatus string

const (
	StatusActive TaskStatus = "in progress"
	StatusCompleted TaskStatus = "completed"
)

type Task struct {
	Id int             `json:"Id"`
	Name string        `json:"Name"`
	Description string `json:"Description"`
	DueDate string     `json:"DueDate"`
	Status TaskStatus  `json:"Status"`
}

func CreateNewTask (id int, name, description, dueDate string, status TaskStatus) *Task {
	return &Task{Id: id, Name: name, Description: description, DueDate: dueDate, Status: status}
}

func (t *Task) PrintTask() {
	fmt.Printf(
		"Task id: %d \n name: %s \n description: %s \n dueDate: %s \n status: %s \n",
		t.Id, t.Name, t.Description, t.DueDate, t.Status,
	)
}

func (t *Task) UpdateTask(name, description, dueDate string, status TaskStatus) {
	if name != "" {
		t.Name = name
	}
	if description != "" {
		t.Description = description
	}
	if dueDate != "" {
		t.DueDate = dueDate
	}
	if status == StatusActive || status == StatusCompleted {
		t.Status = status
	} else {
		fmt.Println("Not a valid status")
	}
}

func UpdateInSlice(tasks []Task, id int, name, description, dueDate string, status TaskStatus) {
	for idx := range tasks {
		if tasks[idx].Id == id {
			var taskToUpdate = &tasks[idx]
			taskToUpdate.UpdateTask(name, description, dueDate, status)
		}
	}
}

func PrintAllTasks(tasks []Task) {
	fmt.Println("#######Print all my tasks bellow:#######")
	 for _, task := range tasks {
		fmt.Println("-----")
		task.PrintTask()
		fmt.Println("-----")
	 }
}

func DeleteFromTasks(tasks []Task, id int) []Task {
	var newTasks []Task 
	for _, task := range tasks {
		if task.Id != id {
			newTasks = append(newTasks, task)
		}
	}
	return newTasks
}

func ToJson(tasks []Task) string {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
	return string(data)
}

func StoreToFile(jsonTasks, fileName string) {
	err := os.WriteFile(fileName, []byte(jsonTasks), 0644)
	if err != nil {
		panic(err)
	}
}

func SearchTask(tasks []Task, id int) Task {
	var task Task
	for idx := range tasks {
		if tasks[idx].Id == id {
			task = tasks[idx]
		}
	}
	return task
}

func ReadFile(fileName string) []Task {
	var tasks []Task
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(string(data)), &tasks)
	return tasks
}