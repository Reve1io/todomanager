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

func (task Task) loadTasks() []Task {

	filepath := "D://go//TodoList/tasks.json"

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Printf("Файл '%s' не существует", filepath)
		return []Task{}
	}

	fileRead, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v", err)
	}

	if len(fileRead) == 0 {
		return []Task{}
	}

	var tasks []Task
	if err := json.Unmarshal(fileRead, &tasks); err != nil {
		return nil
	}

	return tasks
}

func (task Task) add() {

	jsonTask, err := json.MarshalIndent(task, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileWrite := os.WriteFile("tasks.json", jsonTask, 0644)
	if fileWrite != nil {
		fmt.Println("Ошибка при записи файла:", fileWrite)
	}
	fmt.Println("Объект успешно записан")
}

func main() {
	task := Task{ID: 2, Title: os.Args[2], Done: false}
	task.add()
	task.loadTasks()
}
