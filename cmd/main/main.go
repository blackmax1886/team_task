package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type Task struct {
	Id      int
	Name    string
	Content string
}

func main() {
	db, err := sql.Open("mysql", "taskgo:teamtask@tcp(127.0.0.1:3306)/teamtask")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r, err := AddTask(nil, db)
	if err != nil {
		panic(err)
	}
	//fmt.Println(task.Id, task.Name, task.Content)
	fmt.Println(r)

	tasks, err := GetTask(3, db)
	if err != nil {
		panic(err)
	}
	for _, task := range tasks {
		fmt.Println(task.Id, task.Name, task.Content)
	}
}

func AddTask(input *os.File, db *sql.DB) (sql.Result, error) {
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
		return nil, err
	}

	ins, err := db.Prepare("INSERT INTO task (name, content) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	result, err := ins.Exec(task.Name, task.Content)
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func GetTask(limit int, db *sql.DB) ([]*Task, error) {
	const query = "SELECT * FROM task ORDER BY id DESC LIMIT ?"
	rows, err := db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Name, &task.Content)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
