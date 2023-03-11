package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/shresthalucky/go-todo/db"
	"github.com/shresthalucky/go-todo/route"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = db.Setup(os.Getenv("MONGO_URI"), os.Getenv("MONGO_INITDB_DATABASE"))
	if err != nil {
		panic(err)
	}

	r := route.SetupRouters()
	r.Run(":8080")
}
