package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func loadTasks() []Task {
	filepath := "tasks.json"

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
		fmt.Println("Файл пуст")
		return []Task{}
	}

	var tasks []Task
	errors := json.Unmarshal(fileRead, &tasks)
	if errors != nil {
		fmt.Println("Ошибка преобразования в структуру:", errors)
		return []Task{}
	}
	return tasks
}

func add(title string) {
	var newID int

	tasks := loadTasks()
	fmt.Println("Задачи:", tasks)

	if len(tasks) == 0 {
		println("Список задач пуст...")
	}

	if len(tasks) == 0 {
		newID = 1
	}

	newID = tasks[len(tasks)-1].ID + 1

	newTask := Task{
		ID:    newID,
		Title: title,
		Done:  false,
	}

	tasks = append(tasks, newTask)
	fmt.Println("Задача добавлена")

	saveTask(tasks)
}

func saveTask(tasks []Task) {
	jsonTask, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonTask))
	}

	filepath := "tasks.json"

	fileWrite := os.WriteFile(filepath, jsonTask, 0644)
	if fileWrite != nil {
		fmt.Println("Ошибка при записи файла:", fileWrite)
	}
	fmt.Println("Объект успешно записан")
}

func listTasks() {
	tasks := loadTasks()

	for i := 0; i < len(tasks); i++ {
		if tasks[i].Done == true {
			fmt.Println("[✓]", tasks[i].ID, ": ", tasks[i].Title)
		} else {
			fmt.Println("[ ]", tasks[i].ID, ": ", tasks[i].Title)
		}
	}
}

func markDone(id int) {
	tasks := loadTasks()
	found := false

	for i := range tasks {

		if tasks[i].ID == id {
			found = true

			if tasks[i].Done {
				fmt.Println("Задача уже выполнена!")
			} else {
				tasks[i].Done = true
				saveTask(tasks)
				fmt.Println("Статус задачи успешно изменен")
			}
			break
		}
	}

	if !found {
		fmt.Printf("Задачи %d не существует", id)
	}
}

func main() {
	if len(os.Args) < 3 {
		if os.Args[1] == "list" {
			return
		}
		fmt.Println("Использование: введите add /*Название задачи*/")
	}

	switch os.Args[1] {
	case "add":
		add(os.Args[2])
	case "list":
		listTasks()
	case "mark":
		taskID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Ошибка конвертации идентификатора задачи:", err)
		}
		markDone(taskID)
	default:
		fmt.Println("Неизвестная команда, попробуйте add")
	}
}
