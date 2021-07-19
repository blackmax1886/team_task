package task

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Task struct {
	Name    string
	Content string
}

func GetTask(task Task) string {
	return fmt.Sprintf(`%v : %v`, task.Name, task.Content)
}

func AddTask(input *os.File) Task {
	if input == nil {
		input = os.Stdin
	}

	reader := csv.NewReader(input)
	var task Task
	fmt.Println("Enter the name of your task\nname,content =")
	record, err := reader.Read()
	if err != nil {
		panic(err)
	}
	task.Name = record[0]
	task.Content = record[1]

	return task
}
