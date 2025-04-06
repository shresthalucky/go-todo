package main

import (
	"context"
	"time"

	"github.com/shresthalucky/go-todo/todo-service/data"
)

func getCollection() data.Collection {
	return DB.GetCollection("todo")
}

func Create(todo data.Todo) (*data.Todo, error) {
	collection := getCollection()
	t := new(data.Todo)
	t.Init()
	t.Copy(todo)

	objID, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		return nil, err
	}
	return GetById(objID.Hex())
}

func GetAll() (*[]data.Todo, error) {
	collection := getCollection()
	results := new([]data.Todo)
	err := collection.FindAll(context.TODO(), results)
	return results, err
}

func GetById(id string) (*data.Todo, error) {
	collection := getCollection()
	result := new(data.Todo)
	err := collection.FindById(context.TODO(), id, result)
	return result, err
}

func Update(id string, todo data.Todo) (*data.Todo, error) {
	collection := getCollection()
	todo.ID = nil
	todo.CreatedAt = nil
	t := time.Now().UTC()
	todo.UpdatedAt = &t
	_, err := collection.UpdateByIDs(context.TODO(), id, todo)
	if err != nil {
		return nil, err
	}
	return GetById(id)
}
