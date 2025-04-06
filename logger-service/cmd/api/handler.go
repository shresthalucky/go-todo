package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shresthalucky/go-todo/logger-service/data"
	"github.com/shresthalucky/go-todo/logger-service/helper"
)

func getCollection() data.Collection {
	return DB.GetCollection("log")
}

func CreateLog(c *gin.Context) {
	log := new(data.LogEntry)
	if err := c.ShouldBind(log); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	data, err := Create(log)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	helper.SuccessResponse(c, http.StatusCreated, data)
}

func Create(log *data.LogEntry) (*data.LogEntry, error) {
	collection := getCollection()
	log.Init()

	_, err := collection.InsertOne(context.TODO(), log)
	if err != nil {
		return nil, err
	}
	return log, nil
}
