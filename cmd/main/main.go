package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type Task struct {
	Id      int
	Name    string
	Content string
}

func AddTask(input *os.File) (Task, error) {
	var task Task

	if input == nil {
		input = os.Stdin
	}

	buf := make([]byte, 1024)
	n, err := input.Read(buf)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(buf[:n], &task)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func main() {
	db, err := sql.Open("mysql", "taskgo:teamtask@tcp(192.168.100.1:3306)/teamtask")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	task, err := AddTask(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(task.Id, task.Name, task.Content)
}
