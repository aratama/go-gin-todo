package todo

import (
	"context"
	"database/sql"
	"log"

	"aratama.github.com/go-gin-todo/ent"
	"modernc.org/sqlite"
)

func Main() {

	// https://github.com/ent/ent/issues/2460
	sql.Register("sqlite3", &sqlite.Driver{})

	// create ent client
	client, err := ent.Open("sqlite3", "todo.sqlite?_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatalf("Failed open db connection: %v\n", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// start Gin
	InitializeGin(client)
}
