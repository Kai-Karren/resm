package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kai-Karren/resm/managers"
	"github.com/gin-gonic/gin"
)

// Follows Rasa NLG Server API https://rasa.com/docs/rasa/nlg/

type rasaNlgRequest struct {
	Response  string                 `json:"response"`
	Arguments map[string]interface{} `json:"arguments"`
	Tracker   tracker                `json:"tracker"`
	Channel   channel                `json:"channel"`
}

type tracker struct {
	SenderId      string            `json:"sender_id"`
	Slots         map[string]string `json:"slots"`
	LatestMessage latestMessage     `json:"latest_message"`
	Events        []interface{}     `json:"events"`
}

type latestMessage struct {
	MessageId     string                 `json:"message_id"`
	Intent        intent                 `json:"intent"`
	Entities      []interface{}          `json:"entities"`
	Text          string                 `json:"text"`
	Metadata      map[string]interface{} `json:"metadata"`
	IntentRanking []intent               `json:"intent_ranking"`
}

type intent struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	Confidence float32 `json:"confidence"`
}

type channel struct {
	Name string `json:"name"`
}

type rasaNlgResponse struct {
	Text string `json:"text"`
}

type RasaAPI struct {
	ResponseManager managers.StaticResponseManager
}

func (api *RasaAPI) HandleRequest(c *gin.Context) {

	var req rasaNlgRequest

	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(req.Response)

	// fmt.Println(req.Tracker)

	deResponse, err := api.ResponseManager.GetResponse(req.Response)

	if err == nil {
		var res = rasaNlgResponse{
			Text: deResponse,
		}
		c.IndentedJSON(http.StatusOK, res)
	} else {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "500", "message": err.Error()})
	}

}
