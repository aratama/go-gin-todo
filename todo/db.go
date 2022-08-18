package todo

import (
	"context"
	"fmt"

	"log"

	"aratama.github.com/go-gin-todo/ent"
)

const dbFileName = "./todo.sqlite"

func AddTask(client *ent.Client, name string) {
	ctx := context.Background()
	_, err := client.TodoTask.Create().SetName(name).Save(ctx)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	log.Printf("Task %s added\n", name)
}

func RemoveTask(client *ent.Client, id int) {
	ctx := context.Background()
	err := client.TodoTask.DeleteOneID(id).Exec(ctx)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	log.Printf("Task %d deleted\n", id)
}

func GetTodoList(client *ent.Client) []*ent.TodoTask {
	ctx := context.Background()
	todoTaskList, err := client.TodoTask.Query().All(ctx)
	if err != nil {
		log.Printf("%v\b", err)
		var empty []*ent.TodoTask
		return empty
	}
	fmt.Printf("%v\n", todoTaskList)
	return todoTaskList
}
