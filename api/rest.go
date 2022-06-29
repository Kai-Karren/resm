package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kai-Karren/resm/managers"
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

var example_response = response{
	Response: "example response",
}

type SimpleAPI struct {
	ResponseManager managers.StaticResponseManager
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
