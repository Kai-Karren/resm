package managers

import (
	"testing"
)

func TestGetResponse(t *testing.T) {

	var deResponses = make(map[string]string)

	var testResponse = "test, test."

	deResponses["utter_test"] = testResponse

	var responseManager = StaticResponseManager{
		NameToResponse: deResponses,
	}

	response, err := responseManager.GetResponse("utter_test")

	if err != nil {
		t.Fail()
	}

	if response != testResponse {
		t.Fail()
	}

}
