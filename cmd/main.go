package main

import (
	"log"

	"github.com/ocakhasan/mongoapi/internal/controllers"
	"github.com/ocakhasan/mongoapi/internal/repository"
	"github.com/ocakhasan/mongoapi/pkg/database"
	"github.com/ocakhasan/mongoapi/pkg/router"
)

func main() {

	uri := "mongodb://root:pass@localhost:27017"
	mongoDb, err := database.NewMongoDatabase(uri)
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := repository.New(mongoDb)

	controller := controllers.New(repo)

	r := router.Initialize(controller)

	r.Start(":3030")
}
