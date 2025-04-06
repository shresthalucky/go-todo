package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shresthalucky/go-todo/todo-service/data"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	assert := assert.New(t)
	router := SetupRouters()

	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	assert.Nil(err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(http.StatusOK, rr.Code)
	assert.Equal("ping", rr.Body.String())
}

func TestCreateTodo(t *testing.T) {
	assert := assert.New(t)

	_, err := data.Setup("mongodb://mongodb:27017", "go-todo")
	assert.Nil(err)

	router := SetupRouters()

	todo := data.Todo{
		Title: "test title",
		Done:  false,
	}
	b, err := json.Marshal(todo)
	assert.Nil(err)

	req, err := http.NewRequest(http.MethodPost, "/todos", bytes.NewReader(b))
	assert.Nil(err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(http.StatusCreated, rr.Code)
}
