package main

import (
	"github.com/Kai-Karren/resm/api"
	"github.com/Kai-Karren/resm/rasa"
	"github.com/Kai-Karren/resm/storage"
	"github.com/gin-gonic/gin"
)

func main() {

	var responseStorage = storage.NewInMemoryResponseStorage()

	storage.AddResponsesFromJson(&responseStorage, "responses.json")

	var responseGenerator = rasa.NewStaticResponseGenerator(&responseStorage)

	var api = api.NewRasaAPI(&responseGenerator)

	router := gin.Default()
	router.POST("/nlg", api.HandleRequest)

	router.Run("localhost:8080")

}
