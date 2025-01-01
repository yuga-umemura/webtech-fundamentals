package main

import (
	"html/template"
	"net/http"
)

var todoList []string

func handleTodo(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/todo.html")
	t.Execute(w, todoList)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	todo := r.Form.Get("todo")
	todoList = append(todoList, todo)
	http.Redirect(w, r, "/todo", http.StatusSeeOther)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/todo", handleTodo)

	http.HandleFunc("/add", handleAdd)

	http.ListenAndServe(":8080", nil)
}
