package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kai-Karren/resm/responses"
	"github.com/gin-gonic/gin"
)

type request struct {
	Response string            `json:"response"`
	Slots    map[string]string `json:"slots"`
}

type response struct {
	Response string `json:"response"`
	Text     string `json:"text"`
}

type SimpleAPI struct {
	ResponseManager responses.StaticResponseManager
}

func (simpleApi *SimpleAPI) HandleRequest(c *gin.Context) {

	var req request

	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(req.Response)

	fmt.Println(req.Slots)

	deResponse, err := simpleApi.ResponseManager.GetResponse(req.Response)

	deResponse = fillVariablesIfPresent(deResponse, req.Slots)

	if err == nil {
		var res = response{
			Response: "test",
			Text:     deResponse,
		}
		c.IndentedJSON(http.StatusOK, res)
	} else {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "500", "message": err.Error()})
	}

}

func fillVariablesIfPresent(response string, slots map[string]string) string {

	if len(slots) > 0 {

		return responses.FillSlots(response, slots)

	}

	return response

}
