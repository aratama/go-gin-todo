package todo

import (
	"context"
	"log"

	"aratama.github.com/go-gin-todo/ent"
	"github.com/gin-gonic/gin"
)

var client *ent.Client
var router *gin.Engine

func InitizeServer(driverName string, dataSourceName string) *gin.Engine {
	if client == nil || router == nil {

		// create ent client
		client, err := ent.Open(driverName, dataSourceName)
		if err != nil {
			log.Fatalf("Failed open db connection: %v\n", err)
		}

		err = client.Schema.Create(context.Background())
		if err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}

		// start Gin
		router = InitializeGin(client)
	}

	return router
}
