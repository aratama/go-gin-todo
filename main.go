package main

import (
	"database/sql"
	"net/http"

	"aratama.github.com/go-gin-todo/todo"
	"modernc.org/sqlite"
)

func main() {
	// https://github.com/ent/ent/issues/2460
	sql.Register("sqlite3", &sqlite.Driver{})

	router := todo.InitizeServer("sqlite3", "todo.sqlite?_pragma=foreign_keys(1)")
	router.Run("localhost:8080")
}

func ServeForCloudFunctions(w http.ResponseWriter, r *http.Request) {
	router := todo.InitizeServer("sqlite3", "todo.sqlite?_pragma=foreign_keys(1)")
	router.ServeHTTP(w, r)
}
