package storage

import (
	"errors"
	"math/rand"
)

type ResponseStorage interface {
	GetFirstResponse(responseName string) (string, error)
	GetRandomResponse(responseName string) (string, error)
	GetResponses(responseName string) ([]string, error)
	AddResponse(responseName string, response []string)
	DeleteResponse(responseName string)
}

type InMemoryResponseStorage struct {
	responses map[string][]string
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

	if err == nil {
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

func (storage *InMemoryResponseStorage) AddResponse(responseName string, response []string) {

	storage.responses[responseName] = response

}

func (storage *InMemoryResponseStorage) DeleteResponse(responseName string) {

	delete(storage.responses, responseName)

}
