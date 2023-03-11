package test

import (
	"bytes"
	"encoding/json"
	"example/todo/db"
	"example/todo/model"
	"example/todo/route"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	assert := assert.New(t)
	router := route.SetupRouters()

	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	assert.Nil(err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(http.StatusOK, rr.Code)
	assert.Equal("ping", rr.Body.String())
}

func TestCreateTodo(t *testing.T) {
	assert := assert.New(t)

	err := db.Setup("go-todo")
	assert.Nil(err)

	router := route.SetupRouters()

	todo := model.Todo{
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
