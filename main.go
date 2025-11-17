package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func loadTasks() []Task {

	filepath := "D:/go/TodoList/tasks.json"

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Printf("Файл '%s' не существует", filepath)
		return []Task{}
	}

	fileRead, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v", err)
		return []Task{}
	}

	if len(fileRead) == 0 {
		return []Task{}
	}

	var tasks []Task
	if err := json.Unmarshal(fileRead, &tasks); err != nil {
		return []Task{}
	}

	return tasks
}

func (task *Task) add() {
	tasks := loadTasks()

	newID := tasks[len(tasks)-1].ID + 1

	if len(tasks) == 0 {
		newID = 1
	}

	newTask := Task{
		ID:    newID,
		Title: os.Args[2],
		Done:  false,
	}

	tasks = append(tasks, newTask)

	saveTask(tasks)
	fmt.Println("Задача добавлена")
}

func saveTask(tasks []Task) {
	jsonTask, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	filepath := "D:/go/TodoList/tasks.json"

	fileWrite := os.WriteFile(filepath, jsonTask, 0644)
	if fileWrite != nil {
		fmt.Println("Ошибка при записи файла:", fileWrite)
	}
	fmt.Println("Объект успешно записан")
}
