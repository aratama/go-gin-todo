package todo

import (
	"log"
	"strconv"

	"aratama.github.com/go-gin-todo/ent"
	"github.com/gin-gonic/gin"
)

func InitializeGin(client *ent.Client) {

	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.Static("/assets", "./assets")

	router.GET("/", func(c *gin.Context) {
		RenderTaskList(client, c)
	})

	router.GET("/add", func(c *gin.Context) {
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

	router.Run("localhost:8080")
}
