package main

import (
	"database/storage"
	"fmt"
	"os"
)

func main() {

	password := os.Getenv("dbpass")
	connStr := "postgresql://postgres:" + password + "@localhost:5432/todo_list"
	dbase, err := storage.NewStorage(connStr)
	if err != nil {
		fmt.Println(err)
	}

	// Создает новую задачу
	/* 	t := storage.Task{AuthorID: 3, Title: "redis", Content: "create DB redis"}
	   	id, err := dbase.NewTask(t)
	   	if err != nil {
	   		fmt.Println(err)
	   	}
	   	fmt.Printf("Новая задача создана под номером id: %d\n", id) */

	// UpdateTasks обновляет задачу по id
	/* 	title := "clean"
	   	content := "delete all DBs"
	   	err = dbase.UpdateTask(2, title, content)
	   	if err != nil {
	   		fmt.Println(err)
	   	} */

	// DeleteTask удаляет задачу по id
	/* 	err = dbase.DeleteTask(3)
	   	if err != nil {
	   		fmt.Println(err)
	   	} */

	// Возращает список задач по автору
	/* 	taskByAuthor, err := dbase.TasksByAuthor(4)
	   	if err != nil {
	   		fmt.Println(err)
	   	}
	   	for _, val := range taskByAuthor {
	   		fmt.Printf("%v\n", val)
	   	} */

	// Возвращает список всех задач по метке
	/* 	taskByLabel, err := dbase.TasksByLabel("postgres")
	   	if err != nil {
	   		fmt.Println(err)
	   	}
	   	for _, val := range taskByLabel {
	   		fmt.Printf("%v\n", val)
	   	} */

	// Возвращает список всех задач
	tasks, err := dbase.Tasks()
	if err != nil {
		fmt.Println(err)
	}
	for _, val := range tasks {
		fmt.Printf("%v\n", val)
	}
}
