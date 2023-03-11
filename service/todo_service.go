package service

import (
	"context"

	"github.com/shresthalucky/go-todo/db"
	"github.com/shresthalucky/go-todo/model"
	"github.com/shresthalucky/go-todo/util"
)

func getTodoCollection() db.Collection {
	return db.DB.GetCollection("todo")
}

func CreateTodo(todo model.Todo) (*model.Todo, error) {
	collection := getTodoCollection()
	t := new(model.Todo)
	t.Init()
	t.Copy(todo)

	objID, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		return nil, err
	}
	return GetTodoById(objID.Hex())
}

func GetTodos() (*[]model.Todo, error) {
	collection := getTodoCollection()
	results := new([]model.Todo)
	err := collection.FindAll(context.TODO(), results)
	return results, err
}

func GetTodoById(id string) (*model.Todo, error) {
	collection := getTodoCollection()
	result := new(model.Todo)
	err := collection.FindById(context.TODO(), id, result)
	return result, err
}

func UpdateTodo(id string, todo model.Todo) (*model.Todo, error) {
	collection := getTodoCollection()
	todo.ID = nil
	todo.CreatedAt = nil
	todo.UpdatedAt = util.GetCurrentUTCTime()
	_, err := collection.UpdateByIDs(context.TODO(), id, todo)
	if err != nil {
		return nil, err
	}
	return GetTodoById(id)
}
