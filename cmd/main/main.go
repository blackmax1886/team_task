package main

import (
	"encoding/json"
	"fmt"
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
	task, err := AddTask(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(task.Id, task.Name, task.Content)
}
