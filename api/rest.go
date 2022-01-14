package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	Type string `json:"type"`
	// Slots map[string]string `json:"slots"`
	// does not work as expected yet I think
	Slots []struct {
		values map[string]interface{}
	} `json:"slots"`
}

type response struct {
	Type string `json:"response"`
}

var example_response = response{
	Type: "example response",
}

func HandleRequest(c *gin.Context) {

	var req request

	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(req.Type)

	fmt.Println(req.Slots[0].values)

	c.IndentedJSON(http.StatusOK, example_response)
}

// func main() {
// 	router := gin.Default()
// 	router.POST("/request", handleRequest)

// 	router.Run("localhost:8080")
// }
