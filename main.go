package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Kai-Karren/resm/api"
	"github.com/Kai-Karren/resm/responses"
	"github.com/gin-gonic/gin"
)

func main() {

	deResponses := readJsonFile()

	fmt.Println(deResponses["utter_test"])

	printLoadedResponses(deResponses)

	var responseManager = responses.StaticResponseManager{
		NameToResponse: deResponses,
	}

	var api = api.RasaAPI{
		ResponseManager: responseManager,
	}

	router := gin.Default()
	router.POST("/nlg", api.HandleRequest)

	router.Run("localhost:8080")

}

func readJsonFile() map[string]interface{} {

	data, err := os.ReadFile("responses.json")

	check(err)

	var responses map[string]interface{}

	json.Unmarshal([]byte(data), &responses)

	return responses

}

func printLoadedResponses(responses map[string]interface{}) {

	fmt.Println("Loaded", len(responses), "responses")

	for response := range responses {
		fmt.Print(response + " ")
	}

	fmt.Println()

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
