package controller

import (
	"net/http"

	"github.com/shresthalucky/go-todo/helper"
	"github.com/shresthalucky/go-todo/model"
	"github.com/shresthalucky/go-todo/service"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo model.Todo
	if err := c.ShouldBind(&todo); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	data, err := service.CreateTodo(todo)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	helper.SuccessResponse(c, http.StatusCreated, data)
}

func GetTodos(c *gin.Context) {
	data, err := service.GetTodos()
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	helper.SuccessResponse(c, http.StatusOK, data)
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	data, err := service.GetTodoById(id)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	helper.SuccessResponse(c, http.StatusOK, data)
}

func UpdateTodo(c *gin.Context) {
	var todo model.Todo
	if err := c.ShouldBind(&todo); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	id := c.Param("id")
	data, err := service.UpdateTodo(id, todo)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	helper.SuccessResponse(c, http.StatusOK, data)
}
