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

type Tracker struct {
	SenderId      string            `json:"sender_id"`
	Slots         map[string]string `json:"slots"`
	LatestMessage LatestMessage     `json:"latest_message"`
	Events        []interface{}     `json:"events"`
}

type LatestMessage struct {
	MessageId     string                 `json:"message_id"`
	Intent        Intent                 `json:"intent"`
	Entities      []interface{}          `json:"entities"`
	Text          string                 `json:"text"`
	Metadata      map[string]interface{} `json:"metadata"`
	IntentRanking []Intent               `json:"intent_ranking"`
}

type Intent struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	Confidence float32 `json:"confidence"`
}

type Channel struct {
	Name string `json:"name"`
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
