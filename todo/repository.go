package todo

import (
	"context"
	"log"

	"aratama.github.com/go-gin-todo/ent"
)

func AddTask(client *ent.Client, name string) {
	ctx := context.Background()
	_, err := client.TodoTask.Create().SetName(name).Save(ctx)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
}

func RemoveTask(client *ent.Client, id int) {
	ctx := context.Background()
	err := client.TodoTask.DeleteOneID(id).Exec(ctx)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
}

func GetTodoList(client *ent.Client) []*ent.TodoTask {
	ctx := context.Background()
	todoTaskList, err := client.TodoTask.Query().All(ctx)
	if err != nil {
		log.Printf("Error %v\b", err)
		var empty []*ent.TodoTask
		return empty
	}
	return todoTaskList
}
