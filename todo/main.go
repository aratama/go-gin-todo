package todo

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func RenderTaskList(db *sql.DB, c *gin.Context) {
	tasks := GetTodoList(db)
	c.HTML(http.StatusOK, "index.go.tmpl", gin.H{
		"tasks": tasks,
	})
}

func TodoMain() {

	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	InitializeTables(db)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		RenderTaskList(db, c)
	})
	router.POST("/add", func(c *gin.Context) {
		c.Request.ParseForm()
		name := c.Request.PostForm.Get("name")
		if name != "" {
			AddTask(db, name)
		}
		RenderTaskList(db, c)
	})
	router.GET("/delete", func(c *gin.Context) {
		RenderTaskList(db, c)
	})
	router.POST("/delete", func(c *gin.Context) {
		c.Request.ParseForm()
		idString := c.Request.PostForm.Get("id")
		if idString != "" {
			id, err := strconv.Atoi(idString)
			if err != nil {
				log.Fatal(err)
			}
			RemoveTask(db, id)
		}
		RenderTaskList(db, c)
	})
	log.Printf("Running on http://localhost:8080\n")
	router.Run("localhost:8080")
}
