package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kai-Karren/resm/rasa"
	"github.com/gin-gonic/gin"
)

// Follows Rasa NLG Server API https://rasa.com/docs/rasa/nlg/

type RasaAPI struct {
	ResponseGenerator rasa.ResponseGenerator
}

func NewRasaAPI(generator rasa.ResponseGenerator) RasaAPI {
	return RasaAPI{
		ResponseGenerator: generator,
	}
}

func (api *RasaAPI) HandleRequest(c *gin.Context) {

	var request rasa.NlgRequest

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
