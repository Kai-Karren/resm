package api

import (
	"fmt"
	"net/http"

	"github.com/Kai-Karren/resm/managers"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name  string            `json:"name"`
	Type  string            `json:"type"`
	Slots map[string]string `json:"slots"`
}

type response struct {
	Type string `json:"response"`
	Text string `json:"text"`
}

var example_response = response{
	Type: "example response",
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

	fmt.Println(req.Name)

	fmt.Println(req.Type)

	fmt.Println(req.Slots)

	de_responses["example_response"] = "This is an example response."

	de_response, err := response_manager.GetResponse(req.Name)

	if err == nil {
		var res = response{
			Type: "test",
			Text: de_response,
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
