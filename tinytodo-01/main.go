package main

import (
	"html/template"
	"log"
	"net/http"
)

var todoList []string

func handleTodo(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/todo.html")
	t.Execute(w, todoList)
}

func main() {
	todoList = append(todoList, "顔を洗う", "朝食を食べる", "歯を磨く")

	// FileServer(http.Dir("static"))の"static"はファイルシステム上で静的ファイルを探すための基点
	// http.StripPrefix("static", ...)の"static"はhttp.FileServer()に渡すパスから"static"という文字列を取り除く
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/todo", handleTodo)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start:", err)
	}

}
