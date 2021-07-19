package task_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"team_task/cmd/task"
	"testing"
)

func TestGetTask(t *testing.T) {
	sampleTask := task.Task{Name: "sample_name", Content: "sample_content"}
	//output := task.GetTask("name", "content")
	output := task.GetTask(sampleTask)
	want := fmt.Sprintf(`%v : %v`, sampleTask.Name, sampleTask.Content)
	if output != want {
		t.Errorf("error")
	}
}

func TestAddTask(t *testing.T) {
	var sampleTask = task.Task{}
	// create temp file for sample Task data
	input, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer func(input *os.File) {
		err := input.Close()
		if err != nil {

		}
	}(input)

	// write sample Task data to temp
	_, err = io.WriteString(input, "codetest,"+"create AddTask function and test it\n")
	if err != nil {
		t.Fatal(err)
	}

	// 書き込んだ直後なのでReaderがEOFを次に読もうとするので、ファイルの先頭に戻す
	_, err = input.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}

	sampleTask = task.AddTask(input)
	if sampleTask.Name != "codetest" || sampleTask.Content != "create AddTask function and test it" {
		t.Errorf(`unexpected value error: {Name:%v, Content:%v}`, sampleTask.Name, sampleTask.Content)
	}
}
