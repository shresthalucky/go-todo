package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shresthalucky/go-todo/todo-service/data"
	"github.com/shresthalucky/go-todo/todo-service/helper"
)

func CreateTodo(c *gin.Context) {
	var todo data.Todo
	if err := c.ShouldBind(&todo); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	data, err := Create(todo)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	helper.SuccessResponse(c, http.StatusCreated, data)
}

func GetTodos(c *gin.Context) {
	data, err := GetAll()
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	helper.SuccessResponse(c, http.StatusOK, data)
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	data, err := GetById(id)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	helper.SuccessResponse(c, http.StatusOK, data)
}

func UpdateTodo(c *gin.Context) {
	var todo data.Todo
	if err := c.ShouldBind(&todo); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	id := c.Param("id")
	data, err := Update(id, todo)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	helper.SuccessResponse(c, http.StatusOK, data)
}
