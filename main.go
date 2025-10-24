package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "modernc.org/sqlite"
)

type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite", "./data/todos.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		details TEXT
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/todos", handleTodos)
	http.HandleFunc("/todos/", handleTodoByID)

	log.Println("✅ Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := db.Query("SELECT id, title, details FROM todos")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var todos []Todo
		for rows.Next() {
			var t Todo
			rows.Scan(&t.ID, &t.Title, &t.Details)
			todos = append(todos, t)
		}
		json.NewEncoder(w).Encode(todos)

	case http.MethodPost:
		var t Todo
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		if t.Title == "" {
			http.Error(w, "Title is required", http.StatusBadRequest)
			return
		}
		_, err := db.Exec("INSERT INTO todos (title, details) VALUES (?, ?)", t.Title, t.Details)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func handleTodoByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/todos/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		var t Todo
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		_, err = db.Exec("UPDATE todos SET title=?, details=? WHERE id=?", t.Title, t.Details, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		_, err = db.Exec("DELETE FROM todos WHERE id=?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}
