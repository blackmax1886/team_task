package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

type Task struct {
	Id      int
	Name    string
	Content string
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func dbConn() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "taskgo"
	dbPass := "teamtask"
	dbName := "teamtask"
	dbHost := "127.0.0.1:3306"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName)
	if err != nil {
		return nil, err
	}
	return db, err
}

func Index(w http.ResponseWriter, r *http.Request) {
	db, err := dbConn()
	selDB, err := db.Query("SELECT * FROM task ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	task := Task{}
	res := []Task{}
	for selDB.Next() {
		var id int
		var name, content string
		err = selDB.Scan(&id, &name, &content)
		if err != nil {
			panic(err.Error())
		}
		task.Id = id
		task.Name = name
		task.Content = content
		res = append(res, task)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)

	//r, err := AddTask(os.Stdin, db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	////fmt.Println(task.Id, task.Name, task.Content)
	//fmt.Println(r)
	//
	//tasks, err := GetTask(3, db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//ShowTasks(tasks)
	//os.Exit(1)
}

//func AddTask(input *os.File, db *sql.DB) (sql.Result, error) {
//	var task Task
//
//	buf := make([]byte, 1024)
//	n, err := input.Read(buf)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = json.Unmarshal(buf[:n], &task)
//	if err != nil {
//		return nil, err
//	}
//
//	ins, err := db.Prepare("INSERT INTO task (name, content) VALUES(?,?)")
//	if err != nil {
//		log.Fatal(err)
//	}
//	result, err := ins.Exec(task.Name, task.Content)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return result, nil
//}
//
//// GetTask get recent tasks
//func GetTask(limit int, db *sql.DB) ([]*Task, error) {
//	const query = "SELECT * FROM task ORDER BY id DESC LIMIT ?"
//	rows, err := db.Query(query, limit)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var tasks []*Task
//	for rows.Next() {
//		var task Task
//		err := rows.Scan(&task.Id, &task.Name, &task.Content)
//		if err != nil {
//			return nil, err
//		}
//		tasks = append(tasks, &task)
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return tasks, nil
//}
//
//func ShowTasks(tasks []*Task) {
//	for _, task := range tasks {
//		fmt.Println(task.Id, task.Name, task.Content)
//	}
//}
