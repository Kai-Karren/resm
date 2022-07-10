package rasa

import (
	"github.com/Kai-Karren/resm/responses"
	"github.com/Kai-Karren/resm/storage"
)

type ResponseGenerator interface {
	Generate(nlgRequest NlgRequest) (NlgResponse, error)
}

type StaticResponseGenerator struct {
	ResponseStorage storage.ResponseStorage
}

func NewStaticResponseGenerator(responseStorage storage.ResponseStorage) StaticResponseGenerator {
	return StaticResponseGenerator{
		ResponseStorage: responseStorage,
	}
}

func (generator *StaticResponseGenerator) Generate(nlgRequest NlgRequest) (NlgResponse, error) {

	response, err := generator.ResponseStorage.GetRandomResponse(nlgRequest.Response)

	if err != nil {
		return NewRasaNlgResponse(""), err
	}

	response = responses.FillVariablesIfPresent(response, nlgRequest.Tracker.Slots)

	return NewRasaNlgResponse(response), nil

}
