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
	Tracker   map[string]interface{} `json:"tracker"`
	Channel   string                 `json:"channel"`
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

	fmt.Println(req.Tracker)

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
