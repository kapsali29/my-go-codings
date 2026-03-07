package task

import (
	"simple-task-cli/internal/config"
)

func StoreTaskManager(tasks []Task) {
	tasksToJson := ToJson(tasks)
	StoreToFile(tasksToJson, config.FileNameDb)
}

func AddTaskManager(id int, name, description, dueDate, status string, store bool) {
	tasks := ReadFile(config.FileNameDb)
	cliTask := CreateNewTask(id, name, description, dueDate, TaskStatus(status))
	tasks = append(tasks, *cliTask)
	if store {
		StoreTaskManager(tasks)
	}
}

func GetTaskManager(idFlag int, allFlag bool) {
	tasksFromFile := ReadFile(config.FileNameDb)
	if allFlag {
		PrintAllTasks(tasksFromFile)
	}
	if idFlag != 0 {
		resultTask := SearchTask(tasksFromFile, idFlag)
		resultTask.PrintTask()
	}
}

func DelTaskManager(idFlag int) {
	tasksFromFile := ReadFile(config.FileNameDb)
	newTasks := DeleteFromTasks(tasksFromFile, idFlag)
	PrintAllTasks(newTasks)
	StoreTaskManager(newTasks)
}

func UpdateTaskManager(idFlag int, name, status, description, dueDate string) {
	tasksFromFile := ReadFile(config.FileNameDb)
	UpdateInSlice(tasksFromFile, idFlag, name, description, dueDate, TaskStatus(status))
	StoreTaskManager(tasksFromFile)
}