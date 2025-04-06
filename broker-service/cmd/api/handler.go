package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

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

	// Log the incoming request
	logAction(c, fmt.Sprintf("Received broker request for action: %s", br.Action))

	switch br.Action {
	case data.CreateTodoAction:
		CreateTodo(c, br.TodoPayload)
	case data.ListTodosAction:
		ListTodos(c)
	case data.GetTodoByIDAction:
		GetTodoByID(c, br.TodoPayload.ID)
	case data.UpdateTodoAction:
		UpdateTodo(c, br.TodoPayload)
	default:
		errMsg := fmt.Sprintf("Unknown action: %s", br.Action)
		logAction(c, errMsg)
		helper.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf(errMsg))
	}
}

func CreateTodo(c *gin.Context, payload data.TodoPayload) {
	logAction(c, fmt.Sprintf("Creating todo with title: %s", payload.Title))

	jsonData, err := json.Marshal(payload)
	if err != nil {
		logError(c, "Failed to marshal todo payload", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	resp, err := http.Post("http://todo-service:8081/todos", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logError(c, "Failed to connect to todo service", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logError(c, "Failed to read todo service response", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	res := new(helper.Response)
	if err := json.Unmarshal(body, res); err != nil {
		logError(c, "Failed to unmarshal todo service response", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	logAction(c, "Successfully created todo")
	helper.SuccessResponse(c, http.StatusCreated, res)
}

func ListTodos(c *gin.Context) {
	logAction(c, "Listing all todos")

	resp, err := http.Get("http://todo-service:8081/todos")
	if err != nil {
		logError(c, "Failed to connect to todo service", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logError(c, "Failed to read todo service response", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	res := new(helper.Response)
	if err := json.Unmarshal(body, res); err != nil {
		logError(c, "Failed to unmarshal todo service response", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	logAction(c, "Successfully retrieved todo list")
	helper.SuccessResponse(c, http.StatusOK, res)
}

func GetTodoByID(c *gin.Context, todoID string) {
	logAction(c, fmt.Sprintf("Getting todo with ID: %s", todoID))

	resp, err := http.Get(fmt.Sprintf("http://todo-service:8081/todos/%s", todoID))
	if err != nil {
		logError(c, "Failed to connect to todo service", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logError(c, "Failed to read todo service response", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if resp.StatusCode == http.StatusNotFound {
		logAction(c, fmt.Sprintf("Todo with ID %s not found", todoID))
		helper.ErrorResponse(c, http.StatusNotFound, fmt.Errorf("todo with ID %s not found", todoID))
		return
	}

	res := new(helper.Response)
	if err := json.Unmarshal(body, res); err != nil {
		logError(c, "Failed to unmarshal todo service response", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	logAction(c, fmt.Sprintf("Successfully retrieved todo with ID: %s", todoID))
	helper.SuccessResponse(c, http.StatusOK, res)
}

func UpdateTodo(c *gin.Context, payload data.TodoPayload) {
	logAction(c, fmt.Sprintf("Updating todo with ID: %s", payload.ID))

	jsonData, err := json.Marshal(payload)
	if err != nil {
		logError(c, "Failed to marshal todo payload", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://todo-service:8081/todos/%s", payload.ID), bytes.NewBuffer(jsonData))
	if err != nil {
		logError(c, "Failed to create update request", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logError(c, "Failed to connect to todo service", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logError(c, "Failed to read todo service response", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if resp.StatusCode == http.StatusNotFound {
		logAction(c, fmt.Sprintf("Todo with ID %s not found for update", payload.ID))
		helper.ErrorResponse(c, http.StatusNotFound, fmt.Errorf("todo with ID %s not found", payload.ID))
		return
	}

	res := new(helper.Response)
	if err := json.Unmarshal(body, res); err != nil {
		logError(c, "Failed to unmarshal todo service response", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	logAction(c, fmt.Sprintf("Successfully updated todo with ID: %s", payload.ID))
	helper.SuccessResponse(c, resp.StatusCode, res)
}

func logAction(c *gin.Context, message string) {
	sendLog(c, "INFO", message, nil)
}

func logError(c *gin.Context, message string, err error) {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	sendLog(c, "ERROR", message, &errMsg)
}

func sendLog(c *gin.Context, level string, message string, errorDetail *string) {
	logEntry := data.LogEntry{
		Level:       level,
		Message:     message,
		ServiceName: "broker-service",
	}

	if errorDetail != nil {
		logEntry.Error = *errorDetail
	}

	if requestID, exists := c.Get("RequestID"); exists {
		logEntry.RequestID = fmt.Sprintf("%v", requestID)
	}

	jsonData, err := json.Marshal(logEntry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal log entry: %v\n", err)
		return
	}

	// Send log asynchronously to not block the main request flow
	go func() {
		resp, err := http.Post("http://logger-service:8082/logs", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to send log to logger service: %v\n", err)
			return
		}
		defer resp.Body.Close()
	}()
}
