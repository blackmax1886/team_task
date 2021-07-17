package task

import "fmt"

//type Task struct {
//	name string
//	content string
//}

func GetTask(name string, content string) string {
	return fmt.Sprintf(`%v : %v`, name, content)
}
