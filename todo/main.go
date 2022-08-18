package todo

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"aratama.github.com/go-gin-todo/ent"

	_ "github.com/mattn/go-sqlite3"
)

type TemplateTask struct {
	Id        int
	Name      string
	CreatedAt string
}

func RenderTaskList(client *ent.Client, c *gin.Context) {
	tasks := GetTodoList(client)
	fmt.Printf("%v\n", tasks)

	taskList := []TemplateTask{}
	for _, t := range tasks {
		taskList = append(taskList, TemplateTask{
			Id:        t.ID,
			Name:      t.Name,
			CreatedAt: t.CreatedAt.String(),
		})
	}
	c.HTML(http.StatusOK, "index.go.tmpl", gin.H{
		"tasks": taskList,
	})
}

func TodoMain() {

	// initialize Sqlite
	// client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	client, err := ent.Open("sqlite3", "todo.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// InitializeTables(db)

	// initilize Gin

	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	router.GET("/", func(c *gin.Context) {
		RenderTaskList(client, c)
	})
	router.POST("/add", func(c *gin.Context) {
		c.Request.ParseForm()
		name := c.Request.PostForm.Get("name")
		if name != "" {
			AddTask(client, name)
		}
		RenderTaskList(client, c)
	})
	router.GET("/delete", func(c *gin.Context) {
		RenderTaskList(client, c)
	})
	router.POST("/delete", func(c *gin.Context) {
		c.Request.ParseForm()
		idString := c.Request.PostForm.Get("id")
		if idString != "" {
			id, err := strconv.Atoi(idString)
			if err != nil {
				log.Fatal(err)
			}
			RemoveTask(client, id)
		}
		RenderTaskList(client, c)
	})
	log.Printf("Running on http://localhost:8080\n")
	router.Run("localhost:8080")
}
