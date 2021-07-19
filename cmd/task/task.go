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
	fmt.Println("Enter the name of your task\nname =")
	record, err := reader.Read()
	//_, err := fmt.Fscanf(reader, "%s,%s", &task.Name,&task.Content)
	if err != nil {
		panic(err)
	}
	task.Name = record[0]
	task.Content = record[1]
	//fmt.Println("Enter the name & content of your task\ncontent =")
	//
	//_, err = fmt.Fscanln(input, "%s",&task.Content)
	//if err != nil {
	//	panic(err)
	//}

	return task
}
