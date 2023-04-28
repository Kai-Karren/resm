package rasa

import (
	"errors"
	"math/rand"

	"github.com/Kai-Karren/resm/responses"
	"github.com/Kai-Karren/resm/storage"
)

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

	responseOptions, err := generator.staticResponseGenerator.ResponseStorage.GetResponses(nlgRequest.Response)

	if err != nil {
		return NewNlgResponse(""), err
	}

	response, err := generator.responseMemory.SelectUnseenResponse(nlgRequest, responseOptions)

	response = responses.FillVariablesIfPresent(response, nlgRequest.Tracker.Slots)

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
