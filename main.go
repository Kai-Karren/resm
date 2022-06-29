package main

import (
	"github.com/Kai-Karren/resm/api"
	"github.com/Kai-Karren/resm/managers"
	"github.com/gin-gonic/gin"
)

func main() {

	var deResponses = make(map[string]string)

	deResponses["example_response"] = "This is an example response."
	deResponses["another_response"] = "This is another example response."
	deResponses["utter_test"] = "test, test."

	var responseManager = managers.StaticResponseManager{
		Name_to_response: deResponses,
	}

	var simpleApi = api.SimpleAPI{
		ResponseManager: responseManager,
	}

	router := gin.Default()
	router.POST("/request", simpleApi.HandleRequest)

	router.Run("localhost:8080")
}
