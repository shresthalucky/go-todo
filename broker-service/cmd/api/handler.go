package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shresthalucky/go-todo/broker-service/data"
	"github.com/shresthalucky/go-todo/broker-service/helper"
)

func HandleRequest(c *gin.Context) {
	var br data.BrokerRequest
	if err := c.ShouldBind(&br); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	switch br.Action {
	case data.CreateTodoAction:
		CreateTodo(c, br.TodoPayload)
	}
}

func CreateTodo(c *gin.Context, payload data.TodoPayload) {
	jsonData, err := json.Marshal(payload)
	if err != nil {

	}
	resp, err := http.Post("http://todo-service:8081/todos", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	helper.SuccessResponse(c, http.StatusCreated, body)
}
