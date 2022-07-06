package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kai-Karren/resm/responses"
	"github.com/gin-gonic/gin"
)

// Follows Rasa NLG Server API https://rasa.com/docs/rasa/nlg/

type RasaNlgRequest struct {
	Response  string                 `json:"response"`
	Arguments map[string]interface{} `json:"arguments"`
	Tracker   Tracker                `json:"tracker"`
	Channel   Channel                `json:"channel"`
}

type RasaNlgResponse struct {
	Text string `json:"text"`
}

func NewRasaNlgResponse(text string) RasaNlgResponse {
	return RasaNlgResponse{
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
	Generate(nlgRequest RasaNlgRequest) (RasaNlgResponse, error)
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
	ResponseManager responses.StaticResponseManager
}

func NewStaticResponseGenerator(responseManager responses.StaticResponseManager) StaticResponseGenerator {
	return StaticResponseGenerator{
		ResponseManager: responseManager,
	}
}

func (generator *StaticResponseGenerator) Generate(nlgRequest RasaNlgRequest) (RasaNlgResponse, error) {

	response, err := generator.ResponseManager.GetResponse(nlgRequest.Response)

	if err != nil {
		return NewRasaNlgResponse(""), err
	}

	response = fillVariablesIfPresent(response, nlgRequest.Tracker.Slots)

	return NewRasaNlgResponse(response), nil

}
