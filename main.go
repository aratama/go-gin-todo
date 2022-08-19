package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"aratama.github.com/go-gin-todo/todo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"modernc.org/sqlite"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	useSqlite := true

	if useSqlite {
		// https://github.com/ent/ent/issues/2460
		sql.Register("sqlite3", &sqlite.Driver{})
		router := todo.InitizeServer("sqlite3", "todo.sqlite?_pragma=foreign_keys(1)")
		router.Run("localhost:8080")
	} else {
		// NOte: need parseTime=true option in dataSouce
		router := todo.InitizeServer("mysql", os.Getenv("DSN"))
		router.Run("localhost:8080")
	}

}

func ServeForCloudFunctions(w http.ResponseWriter, r *http.Request) {
	router := todo.InitizeServer("sqlite3", "todo.sqlite?_pragma=foreign_keys(1)")
	router.ServeHTTP(w, r)
}
