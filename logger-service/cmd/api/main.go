package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shresthalucky/go-todo/logger-service/data"
)

var DB *data.Database

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db, err := data.Setup(os.Getenv("MONGO_URI"), os.Getenv("MONGO_INITDB_DATABASE"))
	if err != nil {
		panic(err)
	}
	DB = db

	r := SetupRouters()
	r.Run(":8082")
}

func SetupRouters() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	InitTodoRoutes(router)

	return router
}

func InitTodoRoutes(r *gin.Engine) {
	tr := r.Group("/logs")
	{
		tr.POST("", CreateLog)
	}
}
