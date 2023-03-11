package main

import (
	"example/todo/db"
	"example/todo/route"
)

func main() {
	err := db.Setup("go-todo")
	if err != nil {
		panic(err)
	}

	r := route.SetupRouters()
	r.Run(":8080")
}
