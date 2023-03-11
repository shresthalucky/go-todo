package route

import "github.com/gin-gonic/gin"

func SetupRouters() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	InitTodoRoutes(router)

	return router
}
