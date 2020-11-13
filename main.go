package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Todo Type
type Todo struct {
	Name       string `json:"name"`
	IsFinished bool   `json:"isFinished"`
}

func main() {
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/todos", todoHandler)

	http.ListenAndServe(":8080", nil)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllTodos(w, r)
	case "POST":
		postTodo(w, r)
	default:
		fmt.Fprintf(w, "Method not allowed.")
	}
}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{{"todo1", false}, {"todo2", true}}

	// encode json
	json, _ := json.Marshal(todos)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	// 文字列→数値変換
	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))

	// スライス定義
	body := make([]byte, length)
	length, _ = r.Body.Read(body)

	//body →　構造体
	var todo Todo
	json.Unmarshal(body, &todo)

	json, _ := json.Marshal(todo)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
