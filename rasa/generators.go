package rasa

import (
	"errors"

	"github.com/Kai-Karren/resm/responses"
	"github.com/Kai-Karren/resm/storage"
)

type ResponseGenerator interface {
	Generate(nlgRequest NlgRequest) (NlgResponse, error)
	GetHandeledResponses() []string
	HandlesResponse(responseName string) bool
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

func (generator *StaticResponseGenerator) GetHandeledResponses() []string {

	return generator.ResponseStorage.GetAllResponseNames()

}

func (generator *StaticResponseGenerator) HandlesResponse(responseName string) bool {

	return generator.ResponseStorage.HasResponseFor(responseName)

}

type DistributedResponseGenerator struct {
	generators []ResponseGenerator
}

func NewDistributedResponseGenerator() DistributedResponseGenerator {
	return DistributedResponseGenerator{
		generators: []ResponseGenerator{},
	}
}

func (generator *DistributedResponseGenerator) Generate(nlgRequest NlgRequest) (NlgResponse, error) {

	responseName := nlgRequest.Response

	for _, gen := range generator.generators {

		if gen.HandlesResponse(responseName) {
			return gen.Generate(nlgRequest)
		}

	}

	return NewRasaNlgResponse(""), errors.New("no response generator could handle the request")

}

func (generator *DistributedResponseGenerator) GetHandeledResponses() []string {

	combined := []string{}

	for _, g := range generator.generators {
		combined = append(combined, g.GetHandeledResponses()...)
	}

	return combined

}

func (generator *DistributedResponseGenerator) HandlesResponse(responseName string) bool {

	for _, g := range generator.generators {
		if g.HandlesResponse(responseName) {
			return true
		}
	}

	return false

}

func (generator *DistributedResponseGenerator) AddGenerator(responseGenerator ResponseGenerator) {
	generator.generators = append(generator.generators, responseGenerator)
}

type CustomResponseGenerator struct {
	handlers map[string]func(NlgRequest) (NlgResponse, error)
}

func NewCustomResponseGenerator() CustomResponseGenerator {
	return CustomResponseGenerator{
		make(map[string]func(NlgRequest) (NlgResponse, error)),
	}
}

func (generator *CustomResponseGenerator) Generate(nlgRequest NlgRequest) (NlgResponse, error) {

	handler := generator.handlers[nlgRequest.Response]

	return handler(nlgRequest)

}

func (generator *CustomResponseGenerator) GetHandeledResponses() []string {

	combined := []string{}

	for key := range generator.handlers {
		combined = append(combined, key)
	}

	return combined

}

func (generator *CustomResponseGenerator) HandlesResponse(responseName string) bool {

	for key := range generator.handlers {
		if key == responseName {
			return true
		}
	}

	return false

}

func (generator *CustomResponseGenerator) AddHandler(responseName string, handler func(NlgRequest) (NlgResponse, error)) {

	generator.handlers[responseName] = handler

}
