package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shresthalucky/go-todo/db"
	"github.com/shresthalucky/go-todo/model"
	"github.com/shresthalucky/go-todo/route"

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

	err := db.Setup("mongodb://mongodb:27017", "go-todo")
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
