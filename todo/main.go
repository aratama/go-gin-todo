package todo

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func TodoMain() {

	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	Initialize(db)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		tasks := GetTodoList(db)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"tasks": tasks,
		})
	})
	router.Run("localhost:8080")
}
