package storage

import (
	"errors"
	"math/rand"
)

type ResponseStorage interface {
	GetFirstResponse(responseName string) (string, error)
	GetRandomResponse(responseName string) (string, error)
	GetResponses(responseName string) ([]string, error)
	SetResponses(responseName string, response []string)
	AddResponses(responseName string, response []string)
	DeleteResponses(responseName string)
}

type InMemoryResponseStorage struct {
	responses map[string][]string
}

func NewInMemoryResponseStorage() InMemoryResponseStorage {
	return InMemoryResponseStorage{
		responses: make(map[string][]string),
	}
}

func (storage *InMemoryResponseStorage) GetFirstResponse(responseName string) (string, error) {

	response := storage.responses[responseName][0]

	if response == "" {
		return "", errors.New("No response exists for responseName " + responseName)
	}

	return response, nil

}

func (storage *InMemoryResponseStorage) GetRandomResponse(responseName string) (string, error) {

	responses, err := storage.GetResponses(responseName)

	if err != nil {
		return "", err
	}

	randomIndex := rand.Intn(len(responses))

	return responses[randomIndex], nil

}

func (storage *InMemoryResponseStorage) GetResponses(responseName string) ([]string, error) {

	responses := storage.responses[responseName]

	if responses == nil {
		return []string{}, errors.New("No response exists for responseName " + responseName)
	}

	return responses, nil

}

func (storage *InMemoryResponseStorage) SetResponses(responseName string, responses []string) {

	storage.responses[responseName] = responses

}

func (storage *InMemoryResponseStorage) AddResponses(responseName string, responses []string) {

	res, err := storage.GetResponses(responseName)

	if err != nil {
		storage.SetResponses(responseName, responses)
		return
	}

	res = append(res, responses...)

	storage.SetResponses(responseName, res)

}

func (storage *InMemoryResponseStorage) DeleteResponses(responseName string) {

	delete(storage.responses, responseName)

}
