package todo

import (
	"net/http"

	"aratama.github.com/go-gin-todo/ent"
	"github.com/gin-gonic/gin"
)

type TemplateTask struct {
	Id        int
	Name      string
	CreatedAt string
}

func RenderTaskList(client *ent.Client, c *gin.Context) {
	tasks := GetTodoList(client)
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
