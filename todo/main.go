package todo

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

func RenderTaskList(db *sql.DB, c *gin.Context) {
	tasks := GetTodoList(db)
	c.HTML(http.StatusOK, "index.go.tmpl", gin.H{
		"tasks": tasks,
	})
}

func TodoMain() {

	// initialize Sqlite
	db, err := sql.Open("sqlite", dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	InitializeTables(db)

	// initilize Gin

	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
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
