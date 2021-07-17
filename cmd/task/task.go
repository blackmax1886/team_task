package task

import "fmt"

type Task struct {
	Name    string
	Content string
}

func GetTask(task Task) string {
	return fmt.Sprintf(`%v : %v`, task.Name, task.Content)
}
