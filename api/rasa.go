package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kai-Karren/resm/rasa"
	"github.com/Kai-Karren/resm/storage"
	"github.com/gin-gonic/gin"
)

// Follows Rasa NLG Server API https://rasa.com/docs/rasa/nlg/

type RasaNlgRequest struct {
	Response  string                 `json:"response"`
	Arguments map[string]interface{} `json:"arguments"`
	Tracker   rasa.Tracker           `json:"tracker"`
	Channel   rasa.Channel           `json:"channel"`
}

type NlgResponse struct {
	Text string `json:"text"`
}

func NewRasaNlgResponse(text string) NlgResponse {
	return NlgResponse{
		Text: text,
	}
}

type RasaAPI struct {
	ResponseGenerator ResponseGenerator
}

func NewRasaAPI(generator ResponseGenerator) RasaAPI {
	return RasaAPI{
		ResponseGenerator: generator,
	}
}

type ResponseGenerator interface {
	Generate(nlgRequest RasaNlgRequest) (NlgResponse, error)
}

func (api *RasaAPI) HandleRequest(c *gin.Context) {

	var request RasaNlgRequest

	if err := c.BindJSON(&request); err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "500", "message": err.Error()})
		return
	}

	response, err := api.ResponseGenerator.Generate(request)

	if err == nil {
		c.IndentedJSON(http.StatusOK, response)
	} else {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "500", "message": err.Error()})
	}

}

type StaticResponseGenerator struct {
	ResponseStorage storage.ResponseStorage
}

func NewStaticResponseGenerator(responseStorage storage.ResponseStorage) StaticResponseGenerator {
	return StaticResponseGenerator{
		ResponseStorage: responseStorage,
	}
}

func (generator *StaticResponseGenerator) Generate(nlgRequest RasaNlgRequest) (NlgResponse, error) {

	response, err := generator.ResponseStorage.GetRandomResponse(nlgRequest.Response)

	if err != nil {
		return NewRasaNlgResponse(""), err
	}

	response = fillVariablesIfPresent(response, nlgRequest.Tracker.Slots)

	return NewRasaNlgResponse(response), nil

}
