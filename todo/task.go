package todo

import (
	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	Id        int
	Name      string
	CreatedAt string
}
