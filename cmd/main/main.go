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
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT * FROM task ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	task := Task{}
	res := []Task{}
	for rows.Next() {
		var id int
		var name, content string
		err = rows.Scan(&id, &name, &content)
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

func Show(w http.ResponseWriter, r *http.Request) {
	db, err := dbConn()
	if err != nil {
		panic(err.Error())
	}

	nId := r.URL.Query().Get("id")
	rows, err := db.Query("SELECT * FROM Task WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	task := Task{}
	for rows.Next() {
		var id int
		var name, content string
		err = rows.Scan(&id, &name, &content)
		if err != nil {
			panic(err.Error())
		}
		task.Id = id
		task.Name = name
		task.Content = content
	}
	tmpl.ExecuteTemplate(w, "Show", task)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db, err := dbConn()
	if err != nil {
		panic(err.Error())
	}

	if r.Method == "POST" {
		name := r.FormValue("name")
		content := r.FormValue("content")
		insForm, err := db.Prepare("INSERT INTO Task(name, content) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, content)
		log.Println("INSERT: Name: " + name + " | Content: " + content)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db, err := dbConn()
	if err != nil {
		panic(err.Error())
	}

	nId := r.URL.Query().Get("id")
	rows, err := db.Query("SELECT * FROM Task WHERE  id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	task := Task{}
	for rows.Next() {
		var id int
		var name, content string
		err = rows.Scan(&id, &name, &content)
		if err != nil {
			panic(err.Error())
		}
		task.Id = id
		task.Name = name
		task.Content = content
	}
	tmpl.ExecuteTemplate(w, "Edit", task)
	defer db.Close()
}

func Update(w http.ResponseWriter, r *http.Request) {
	db, err := dbConn()
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "POST" {
		name := r.FormValue("name")
		content := r.FormValue("content")
		id := r.FormValue("taskid")
		insForm, err := db.Prepare("UPDATE Task SET name=?, content=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, content, id)
		log.Println("UPDATE: Name: " + name + " | Content: " + content)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db, err := dbConn()
	if err != nil {
		panic(err.Error())
	}
	nId := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Task WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(nId)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
