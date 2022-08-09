package main

import (
  "log"
  "net/http"
  "database/sql"
  "text/template"
  _ "github.com/go-sql-driver/mysql"
)

type ToDoList struct {
  ID int
  Assignee string
  Assigner string
  Task string
  Start string
  Finish string
  Status string
}

func dbConn() (db *sql.DB) {
  dbDriver := "mysql"
  dbUser := "root"
  dbPass := ""
  dbName := "GoDB"
  db,err := sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)
  if err != nil {panic(err.Error())}
  return db
}

var tmpl = template.Must(template.ParseGlob("views/*"))

func Index(w http.ResponseWriter, r *http.Request) {
  db := dbConn()
  selDB,err := db.Query("SELECT * FROM ToDoList ORDER BY ID DESC")
  if err != nil {panic(err.Error())}
  tdl := ToDoList{}
  res := []ToDoList{}
  for selDB.Next() {
	var id int
	var assignee,assigner,task,start,finish,status string
	err = selDB.Scan(&id,&assignee,&assigner,&task,&start,&finish,&status)
	if err != nil {panic(err.Error())}
	tdl.ID = id
	tdl.Assignee = assignee
	tdl.Assigner = assigner
	tdl.Task = task
	tdl.Start = start
	tdl.Finish = finish
	tdl.Status = status
	res = append(res,tdl)
  }
  tmpl.ExecuteTemplate(w,"Index",res)
  defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
  db := dbConn()
  nId := r.URL.Query().Get("id")
  selDB,err := db.Query("SELECT * FROM ToDoList WHERE ID=?",nId)
  if err != nil {panic(err.Error())}
  tdl := ToDoList{}
  for selDB.Next() {
	var id int
	var assignee,assigner,task,start,finish,status string
	err = selDB.Scan(&id,&assignee,&assigner,&task,&start,&finish,&status)
	if err != nil {panic(err.Error())}
	tdl.ID = id
	tdl.Assignee = assignee
	tdl.Assigner = assigner
	tdl.Task = task
	tdl.Start = start
	tdl.Finish = finish
	tdl.Status = status
  }
  tmpl.ExecuteTemplate(w,"Show",tdl)
  defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
  tmpl.ExecuteTemplate(w,"New",nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
  db := dbConn()
  nId := r.URL.Query().Get("id")
  selDB,err := db.Query("SELECT * FROM ToDoList WHERE ID=?",nId)
  if err != nil {panic(err.Error())}
  tdl := ToDoList{}
  for selDB.Next() {
	var id int
	var assignee,assigner,task,start,finish,status string
	err = selDB.Scan(&id,&assignee,&assigner,&task,&start,&finish,&status)
	if err != nil {panic(err.Error())}
	tdl.ID = id
	tdl.Assignee = assignee
	tdl.Assigner = assigner
	tdl.Task = task
	tdl.Start = start
	tdl.Finish = finish
	tdl.Status = status
  }
  tmpl.ExecuteTemplate(w,"Edit",tdl)
  defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
  db := dbConn()
  if r.Method == "POST" {
	assignee := r.FormValue("assignee")
	assigner := r.FormValue("assigner")
	task := r.FormValue("task")
	start := r.FormValue("start")
	finish := r.FormValue("finish")
	status := r.FormValue("status")
	insForm,err := db.Prepare("INSERT INTO ToDoList (assignee,assigner,task,start,finish,status) VALUES (?,?,?,?,?,?)")
	if err != nil {panic(err.Error())}
	insForm.Exec(assignee,assigner,task,start,finish,status)
	log.Println("INSERT - Assignee : "+assignee+" | Assigner : "+assigner+" | Task : "+task+" | Start : "+start+" | Finish : "+finish+" | Status : "+status)
  }
  defer db.Close()
  http.Redirect(w,r,"/",301)
}

func Update(w http.ResponseWriter, r * http.Request) {
  db := dbConn()
  if r.Method == "POST" {
	id := r.FormValue("uid")
	assignee := r.FormValue("assignee")
	assigner := r.FormValue("assigner")
	task := r.FormValue("task")
	start := r.FormValue("start")
	finish := r.FormValue("finish")
	status := r.FormValue("status")
	insForm,err := db.Prepare("UPDATE ToDoList SET Assignee=?,Assigner=?,Task=?,Start=?,Finish=?,Status=? WHERE ID=?")
	if err != nil {panic(err.Error())}
	insForm.Exec(assignee,assigner,task,start,finish,status,id)
	log.Println("UPDATE - Assignee : "+assignee+" | Assigner : "+assigner+" | Task : "+task+" | Start : "+start+" | Finish : "+finish+" | Status : "+status)
  }
  defer db.Close()
  http.Redirect(w,r,"/",301)
}

func Delete(w http.ResponseWriter, r * http.Request) {
  db := dbConn()
  tdl := r.URL.Query().Get("id")
  delForm,err := db.Prepare("DELETE FROM ToDoList WHERE ID=?")
  if err != nil {panic(err.Error())}
  delForm.Exec(tdl)
  log.Println("DELETE - ID:"+tdl)
  defer db.Close()
  http.Redirect(w,r,"/",301)
}

func main() {
  log.Println("Server Started On http://localhost:8080")
  http.HandleFunc("/",Index)
  http.HandleFunc("/show",Show)
  http.HandleFunc("/new",New)
  http.HandleFunc("/edit",Edit)
  http.HandleFunc("/insert",Insert)
  http.HandleFunc("/update",Update)
  http.HandleFunc("/delete",Delete)
  http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("assets"))))
  http.ListenAndServe("localhost:8080",nil)
}
