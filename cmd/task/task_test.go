package task_test

import (
	"fmt"
	"team_task/cmd/task"
	"testing"
)

func TestGetTask(t *testing.T) {
	sample_task := task.Task{Name: "sample_name", Content: "sample_content"}
	//output := task.GetTask("name", "content")
	output := task.GetTask(sample_task)
	want := fmt.Sprintf(`%v : %v`, sample_task.Name, sample_task.Content)
	if output != want {
		t.Errorf("error")
	}
}
