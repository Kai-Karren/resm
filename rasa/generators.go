package rasa

import (
	"errors"
	"math/rand"

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
		return NewNlgResponse(""), err
	}

	response = responses.FillVariablesIfPresent(response, nlgRequest.Tracker.Slots)

	return NewNlgResponse(response), nil

}

func (generator *StaticResponseGenerator) GetHandeledResponses() []string {

	return generator.ResponseStorage.GetAllResponseNames()

}

func (generator *StaticResponseGenerator) HandlesResponse(responseName string) bool {

	return generator.ResponseStorage.HasResponseFor(responseName)

}

type StaticResponseGeneratorWithMemory struct {
	staticResponseGenerator StaticResponseGenerator
	responseMemory          ResponseMemory
}

func NewStaticResponseGeneratorWithMemory(responseStorage storage.ResponseStorage) StaticResponseGeneratorWithMemory {
	return StaticResponseGeneratorWithMemory{
		staticResponseGenerator: NewStaticResponseGenerator(responseStorage),
		responseMemory:          NewResponseMemory(),
	}
}

type ResponseMemory struct {
	usersToResponses map[string]ResponsesForUser
}

func NewResponseMemory() ResponseMemory {
	return ResponseMemory{
		usersToResponses: make(map[string]ResponsesForUser),
	}
}

func (responseMemory *ResponseMemory) SelectUnseenResponse(nlgRequest NlgRequest, responses []string) (string, error) {

	if responseMemory.HasResponsesForUserId(nlgRequest.Tracker.SenderId) {

		previousResponses := responseMemory.usersToResponses[nlgRequest.Tracker.SenderId]

		if previousResponses.GetLength(nlgRequest.Response) == len(responses) {

			randomIndex := rand.Intn(len(responses))

			response := responses[randomIndex]

			previousResponses := []string{response}

			responseMemory.usersToResponses[nlgRequest.Tracker.SenderId].responseNamesToSelectedResponses[nlgRequest.Response] = previousResponses

			return response, nil

		} else {

			if previousResponses.HasResponses(nlgRequest.Response) {

				response, err := previousResponses.SelectUnseenResponse(nlgRequest.Response, responses)

				if err != nil {
					return "", err
				}

				responseMemory.usersToResponses[nlgRequest.Tracker.SenderId].responseNamesToSelectedResponses[nlgRequest.Response] = append(responseMemory.usersToResponses[nlgRequest.Tracker.SenderId].responseNamesToSelectedResponses[nlgRequest.Response], response)

				return response, nil

			} else {

				randomIndex := rand.Intn(len(responses))

				response := responses[randomIndex]

				previousResponses := []string{response}

				responseMemory.usersToResponses[nlgRequest.Tracker.SenderId].responseNamesToSelectedResponses[nlgRequest.Response] = previousResponses

				return response, nil

			}

		}

	} else {

		randomIndex := rand.Intn(len(responses))

		response := responses[randomIndex]

		responseMemory.usersToResponses[nlgRequest.Tracker.SenderId] = NewResponsesForUser(nlgRequest.Tracker.SenderId)

		previousResponses := []string{response}

		responseMemory.usersToResponses[nlgRequest.Tracker.SenderId].responseNamesToSelectedResponses[nlgRequest.Response] = previousResponses

		return response, nil

	}

}

func (responseMemory *ResponseMemory) HasResponsesForUserId(userId string) bool {

	_, ok := responseMemory.usersToResponses[userId]

	return ok

}

type ResponsesForUser struct {
	UserId                           string
	responseNamesToSelectedResponses map[string][]string
}

func NewResponsesForUser(userId string) ResponsesForUser {
	return ResponsesForUser{
		UserId:                           userId,
		responseNamesToSelectedResponses: make(map[string][]string),
	}
}

func (responsesForUser *ResponsesForUser) SelectUnseenResponse(responseName string, responses []string) (string, error) {

	previousResponses := responsesForUser.responseNamesToSelectedResponses[responseName]

	for i := 0; i < len(responses); i++ {

		response := responses[i]

		contains := false

		for j := 0; j < len(previousResponses); j++ {

			previousResponseAtIndex := previousResponses[j]

			if response == previousResponseAtIndex {
				contains = true
			}

		}

		if !contains {
			return response, nil
		}

	}

	return "", errors.New("no unseen response could be found")

}

func (responsesForUser *ResponsesForUser) HasResponses(responseName string) bool {

	_, ok := responsesForUser.responseNamesToSelectedResponses[responseName]

	return ok

}

func (responsesForUser *ResponsesForUser) GetLength(responseName string) int {

	return len(responsesForUser.responseNamesToSelectedResponses[responseName])

}

func (generator *StaticResponseGeneratorWithMemory) Generate(nlgRequest NlgRequest) (NlgResponse, error) {

	responses, err := generator.staticResponseGenerator.ResponseStorage.GetResponses(nlgRequest.Response)

	if err != nil {
		return NewNlgResponse(""), err
	}

	response, err := generator.responseMemory.SelectUnseenResponse(nlgRequest, responses)

	if err != nil {
		return NewNlgResponse(""), err
	}

	return NewNlgResponse(response), nil

}

func (generator *StaticResponseGeneratorWithMemory) GetHandeledResponses() []string {

	return generator.staticResponseGenerator.ResponseStorage.GetAllResponseNames()

}

func (generator *StaticResponseGeneratorWithMemory) HandlesResponse(responseName string) bool {

	return generator.staticResponseGenerator.ResponseStorage.HasResponseFor(responseName)

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

	return NewNlgResponse(""), errors.New("no response generator could handle the request")

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

// Returns a static error message that something went wrong. Should only be used called as last option.
// Agrees to handle all reponses!
type DefaultFallbackResponseGenerator struct {
	DefaultFallbackMessage string
}

func NewDefaultFallbackGenerator(defaultFallbackMessage string) DefaultFallbackResponseGenerator {
	return DefaultFallbackResponseGenerator{
		DefaultFallbackMessage: defaultFallbackMessage,
	}
}

func (generator *DefaultFallbackResponseGenerator) Generate(nlgRequest NlgRequest) (NlgResponse, error) {

	return NewNlgResponse(generator.DefaultFallbackMessage), nil

}

func (generator *DefaultFallbackResponseGenerator) GetHandeledResponses() []string {

	return []string{}

}

func (generator *DefaultFallbackResponseGenerator) HandlesResponse(responseName string) bool {

	// Agrees to handle all responses!
	return true

}
