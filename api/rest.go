package api

import (
	"fmt"
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

var de_responses = make(map[string]string)

var response_manager = managers.StaticResponseManager{
	Name_to_response: de_responses,
}

func HandleRequest(c *gin.Context) {

	var req request

	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(req.Response)

	fmt.Println(req.Slots)

	de_responses["example_response"] = "This is an example response."

	de_response, err := response_manager.GetResponse(req.Response)

	if err == nil {
		var res = response{
			Response: "test",
			Text:     de_response,
		}
		c.IndentedJSON(http.StatusOK, res)
	} else {
		c.IndentedJSON(http.StatusInternalServerError, example_response)
	}

}

// func main() {
// 	router := gin.Default()
// 	router.POST("/request", handleRequest)

// 	router.Run("localhost:8080")
// }
