package test

import (
	"bytes"
	"encoding/json"
	"example/todo/controller"
	"example/todo/db"
	"example/todo/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	err := db.Setup("go-todo")
	if err != nil {
		t.Errorf(err.Error())
	}

	todo := model.Todo{
		Title: "test title",
		Done:  false,
	}

	b, _ := json.Marshal(todo)

	req, err := http.NewRequest(http.MethodPost, "/todos", bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(controller.CreateTodo)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Wrong http response status code %d", rr.Code)
	}
}

// func xTestCreateTodos(t *testing.T) {
// 	err := db.Setup("go-todo")

// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}

// 	todos := []interface{}{
// 		model.Todo{
// 			ID:    primitive.NewObjectID(),
// 			Title: "title123",
// 			Done:  false,
// 		},
// 		model.Todo{
// 			ID:    primitive.NewObjectID(),
// 			Title: "title123",
// 			Done:  false,
// 		},
// 	}

// 	err = service.CreateTodos(todos)

// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}

// }
