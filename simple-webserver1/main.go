package main

import (
	"fmt"
	"log"
	"net/http"
)

// リクエストハンドラ（HTTPリクエストを処理する関数）
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Web Application!")
}

func main() {
	// HTTPリクエストを登録
	http.HandleFunc("/", hello)
	// HTTPサーバを起動
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start:", err)
	}
}
