package route

import (
	"example/todo/controller"

	"github.com/gorilla/mux"
)

func InitTodoRoutes(r *mux.Router) {
	todoRouter := r.PathPrefix("/todos").Subrouter()

	todoRouter.HandleFunc("/", controller.GetTodos).Methods("GET")
	todoRouter.HandleFunc("/{id}", controller.GetTodo).Methods("GET")
	todoRouter.HandleFunc("/", controller.CreateTodo).Methods("POST")
	todoRouter.HandleFunc("/{id}", controller.UpdateTodo).Methods("PUT")
}
