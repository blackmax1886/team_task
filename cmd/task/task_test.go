package task_test

import (
	"team_task/cmd/task"
	"testing"
)

func TestGetTask(t *testing.T) {
	//task := task.Task{"test_name", "test_content"}
	output := task.GetTask("name", "content")
	if output != "name : content" {
		t.Errorf("error")
	}

}
