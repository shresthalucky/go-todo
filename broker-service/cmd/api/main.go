package main

import "github.com/gin-gonic/gin"

func main() {
	r := SetupRouters()
	r.Run(":8080")
}

func SetupRouters() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	InitBrokerRoutes(router)

	return router
}

func InitBrokerRoutes(r *gin.Engine) {
	br := r.Group("/broker")
	{
		br.GET("", HandleRequest)
	}
}
