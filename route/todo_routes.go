package route

import (
	"github.com/shresthalucky/go-todo/controller"

	"github.com/gin-gonic/gin"
)

func InitTodoRoutes(r *gin.Engine) {
	tr := r.Group("/todos")
	{
		tr.GET("", controller.GetTodos)
		tr.GET("/:id", controller.GetTodo)
		tr.POST("", controller.CreateTodo)
		tr.PUT("/:id", controller.UpdateTodo)
	}
}
