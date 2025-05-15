package tasks

import (
	"errors"
	"fmt"
	"github.com/k0kubun/pp"
	"log"
	"time"
)

//type TaskStorage interface {
//
//}

type TaskStorages struct {
	currentTasks   []Task
	completedTasks []Task
}

type Task struct {
	title       string
	text        string
	isCompleted bool
	createdAt   int64 //unix
	completedAt int64 //unix
}

func CreateStorages() TaskStorages {
	return TaskStorages{}
}

func (ts *TaskStorages) AddTask(title, text string) error {
	log.Println("Adding task", title, text)
	if len(title) == 0 || len(text) == 0 {
		return errors.New("incorrect values")
		log.Println("Incorrect values", title, text)
	}

	newTask := Task{
		title:       title,
		text:        text,
		isCompleted: false,
		createdAt:   time.Now().Unix(),
		completedAt: 0,
	}
	ts.currentTasks = append(ts.currentTasks, newTask)
	fmt.Printf("Task %s added\n", title)
	log.Println("Adding task", newTask.title, newTask.text)
	return nil
}

func (ts *TaskStorages) ListTasks() error {
	log.Println("Listing tasks")
	hasCurrent := len(ts.currentTasks) > 0
	hasCompleted := len(ts.completedTasks) > 0

	if !hasCurrent && !hasCompleted {
		return errors.New("no tasks found")
	}

	if hasCurrent {
		fmt.Println("Current tasks:")
		for _, task := range ts.currentTasks {
			pp.Printf("Title: %s\nText: %s\n\n", task.title, task.text)
		}
	}

	if hasCompleted {
		fmt.Println("Completed tasks:")
		for _, task := range ts.completedTasks {
			duration := time.Duration(task.completedAt) * time.Second
			formattedDuration := fmt.Sprintf("%02d:%02d:%02d",
				int(duration.Hours()*-1),
				int(duration.Minutes()*-1)%60,
				int(duration.Seconds()*-1)%60)

			pp.Printf("Title: %s\nText: %s\nTime to complete: %s\n\n", task.title, task.text, formattedDuration)
		}
	}

	return nil
}

func (ts *TaskStorages) CompleteTask(title string) error {

	log.Println("Complete task", title)
	index, ok := ts.getTask(title)
	if !ok {
		return errors.New("task not found")
		log.Println("Task not found", title)
	}

	task := ts.currentTasks[index]

	task.isCompleted = true
	task.completedAt = task.createdAt - time.Now().Unix()

	//deleting
	ts.currentTasks = append(ts.currentTasks[:index], ts.currentTasks[index+1:]...)

	ts.completedTasks = append(ts.completedTasks, task)
	fmt.Printf("Task %s completed\n", title)
	log.Println("Complete task", title)
	return nil
}

func (ts *TaskStorages) DeleteTask(title string) error {
	log.Println("Deleting task", title)
	index, ok := ts.getTask(title)
	if !ok {
		return errors.New("task not found")
		log.Println("Task not found", title)
	}
	ts.currentTasks = append(ts.currentTasks[:index], ts.currentTasks[index+1:]...)
	fmt.Println("Deleted task:", title)
	log.Println("Deleted task:", title)
	return nil
}

func (ts *TaskStorages) getTask(title string) (int, bool) {
	if len(ts.currentTasks) == 0 {
		return -1, false
	}
	for index, task := range ts.currentTasks {
		if task.title == title {
			return index, true
		}
	}
	return -1, false
}
