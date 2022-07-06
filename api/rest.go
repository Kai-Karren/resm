package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kai-Karren/resm/responses"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Response string            `json:"response"`
	Slots    map[string]string `json:"slots"`
}

type Response struct {
	Response string `json:"response"`
	Text     string `json:"text"`
}

type SimpleAPI struct {
	ResponseManager responses.StaticResponseManager
}

func (simpleApi *SimpleAPI) HandleRequest(c *gin.Context) {

	var req Request

	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(req.Response)

	fmt.Println(req.Slots)

	deResponse, err := simpleApi.ResponseManager.GetResponse(req.Response)

	deResponse = fillVariablesIfPresent(deResponse, req.Slots)

	if err == nil {
		var res = Response{
			Response: "test",
			Text:     deResponse,
		}
		c.IndentedJSON(http.StatusOK, res)
	} else {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "500", "message": err.Error()})
	}

}

func fillVariablesIfPresent(Response string, slots map[string]string) string {

	if len(slots) > 0 {

		return responses.FillSlots(Response, slots)

	}

	return Response

}
