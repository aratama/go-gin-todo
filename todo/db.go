package todo

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const dbFileName = "./todo.sqlite"

func InitializeTables(db *sql.DB) {
	sqlStmt := `
	create table if not exists todo (
		id integer not null primary key AUTOINCREMENT, 
		name text,
		created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
    	updated_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func AddTask(db *sql.DB, name string) {
	_, err := db.Exec("insert into todo(name) values(?)", name)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Task %s added\n", name)
}

func RemoveTask(db *sql.DB, id int) {
	_, err := db.Exec("delete from todo where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Task %d deleted\n", id)
}

func GetTodoList(db *sql.DB) []Task {
	rows, err := db.Query("select id, name, created_at from todo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var taskList []Task
	for rows.Next() {
		var id int
		var name string
		var created_at string
		err = rows.Scan(&id, &name, &created_at)
		if err != nil {
			log.Fatal(err)
		}
		taskList = append(taskList, Task{id, name, created_at})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return taskList
}
