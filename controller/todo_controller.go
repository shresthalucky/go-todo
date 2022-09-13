package controller

import (
	"example/todo/helper"
	"example/todo/model"
	"example/todo/service"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	err := helper.DecodeRequestBody(r, &todo)
	if err != nil {
		helper.ErrorResponse(w, http.StatusBadRequest, err)
	}

	data, err := service.CreateTodo(todo)
	if err != nil {
		helper.ErrorResponse(w, http.StatusInternalServerError, err)
	}

	helper.SuccessResponse(w, data)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetTodos()
	if err != nil {
		helper.ErrorResponse(w, http.StatusInternalServerError, err)
	}

	helper.SuccessResponse(w, data)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data, err := service.GetTodoById(vars["id"])
	if err != nil {
		helper.ErrorResponse(w, http.StatusInternalServerError, err)
	}

	helper.SuccessResponse(w, data)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	vars := mux.Vars(r)
	err := helper.DecodeRequestBody(r, &todo)
	if err != nil {
		helper.ErrorResponse(w, http.StatusBadRequest, err)
	}

	data, err := service.UpdateTodo(vars["id"], todo)
	if err != nil {
		helper.ErrorResponse(w, http.StatusInternalServerError, err)
	}

	helper.SuccessResponse(w, data)
}
