package task

import (
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

	var name, content string
	fmt.Println("Enter the name & content of your task\nname content =")
	_, err := fmt.Fscanf(input, "%s %s", &name, &content)
	if err != nil {
		panic(err)
	}

	return Task{Name: name, Content: content}
}
