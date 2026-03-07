package main

import (
	"flag"
	"fmt"
	"os"
	"simple-task-cli/internal/task"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	idFlag := addCmd.Int("Id", 0, "Task ID")
	nameFlag := addCmd.String("Name", "", "Task Name")
	descriptionFlag := addCmd.String("Description", "", "Task Description")
	dueDateFlag := addCmd.String("DueDate", "", "Task Due Date")
	status := addCmd.String("Status", "in-progress", "Task Status")
	store := addCmd.Bool("store", false, "store this new task")

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	idGetFlag := getCmd.Int("Id", 0, "Task ID")
	allFlag := getCmd.Bool("all", false, "Get all tasks stored")

	delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	idFlagDel := delCmd.Int("Id", 0, "Task ID")

	updCmd := flag.NewFlagSet("update", flag.ExitOnError)
	idUpdFlag := updCmd.Int("Id", 0, "Task ID")
	nameUpdFlag := updCmd.String("Name", "", "Task Name")
	statusUpd := updCmd.String("Status", "", "Task Status")
	descriptionFlagUpd := updCmd.String("Description", "", "Task Description")
	dueDateFlagUpd := updCmd.String("DueDate", "", "Task Due Date")


	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		task.AddTaskManager(*idFlag, *nameFlag, *descriptionFlag, *dueDateFlag, *status, *store)
	case "get":
		getCmd.Parse(os.Args[2:])
		task.GetTaskManager(*idGetFlag, *allFlag)
	case "delete":
		delCmd.Parse(os.Args[2:])
		task.DelTaskManager(*idFlagDel)
	case "update":
		updCmd.Parse(os.Args[2:])
		task.UpdateTaskManager(*idUpdFlag, *nameUpdFlag, *statusUpd, *descriptionFlagUpd, *dueDateFlagUpd)
	default:
		fmt.Println("No proper command selected")
	}
}
