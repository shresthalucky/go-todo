package main

import (
	"example/todo/db"
	"example/todo/route"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	err := db.Setup("go-todo")
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	route.InitTodoRoutes(router)

	handler := cors.Default().Handler(router)
	http.ListenAndServe(":8080", handler)
}
