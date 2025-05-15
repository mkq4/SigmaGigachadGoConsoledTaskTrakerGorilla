package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"todoList/logger"
	"todoList/tasks"
)

func main() {

	logger.LoggerSetup()
	log.Println("Start")

	scanner := bufio.NewScanner(os.Stdin)
	storages := tasks.CreateStorages()
	fmt.Println("Todo list")
	fmt.Println("Commands: help, add, list, del, done, exit")

	for {
		fmt.Print("\n> ")
		if ok := scanner.Scan(); !ok {
			fmt.Println("Scanner Error")
			break
		}

		rawText := scanner.Text() // scanner

		log.Println("User input:", rawText)

		parts := strings.SplitN(rawText, " ", 2)
		command := parts[0]
		var args string // user args
		if len(parts) > 1 {
			args = parts[1]
			log.Println("Saving args:", args)
		}

		switch command { // TODO: handler commands
		case "help":
			helpCommand()
		case "exit":
			log.Println("Exit")
			fmt.Println("Bye!")
			os.Exit(0)
		case "list":
			err := storages.ListTasks()
			if err != nil {
				log.Println(err)
				fmt.Println(err)
			}
		case "add":
			data := strings.SplitN(args, " ", 2)
			if len(data) < 2 {
				log.Println("Invalid arguments", data)
				fmt.Println("Usage: add <title> <text>")
				continue
			}
			title, text := data[0], data[1]
			err := storages.AddTask(title, text)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "done":
			err := storages.CompleteTask(args)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
				continue
			}
		case "del":
			err := storages.DeleteTask(args)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
				continue
			}
		}
	}
}

func helpCommand() {
	log.Println("Help Command")
	fmt.Println("\nAvailable commands:")
	fmt.Println("help - get information about all commands")
	fmt.Println("add {taskTitle} {taskText} - create new task | {taskTitle} can be only 1 word")
	fmt.Println("list - get information about all tasks")
	fmt.Println("del {taskTitle} - delete task")
	fmt.Println("done {tasksTitle} - completed task")
	fmt.Println("exit - exit program")
}
