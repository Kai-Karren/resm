package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Kai-Karren/resm/api"
	"github.com/Kai-Karren/resm/managers"
	"github.com/gin-gonic/gin"
)

func main() {

	// var deResponses = make(map[string]string)

	// deResponses["example_response"] = "This is an example response."
	// deResponses["another_response"] = "This is another example response."
	// deResponses["utter_test"] = "test, test."

	deResponses := readJsonFile()

	fmt.Println(deResponses["utter_test"])

	printLoadedResponses(deResponses)

	var responseManager = managers.StaticResponseManager{
		NameToResponse: deResponses,
	}

	// var simpleApi = api.SimpleAPI{
	// 	ResponseManager: responseManager,
	// }

	var api = api.RasaAPI{
		ResponseManager: responseManager,
	}

	router := gin.Default()
	router.POST("/nlg", api.HandleRequest)

	router.Run("localhost:8080")

}

func readJsonFile() map[string]string {

	data, err := os.ReadFile("responses.json")

	check(err)

	var responses map[string]string

	json.Unmarshal([]byte(data), &responses)

	return responses

}

func printLoadedResponses(responses map[string]string) {

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
