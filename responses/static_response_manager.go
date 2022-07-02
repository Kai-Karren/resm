package responses

import (
	"errors"
	"math/rand"
)

type StaticResponseManager struct {
	NameToResponse map[string]interface{}
}

func (s *StaticResponseManager) GetResponse(name string) (string, error) {

	if response, ok := s.NameToResponse[name]; ok {

		switch r := response.(type) {
		case string:
			return r, nil
		case []interface{}:
			return r[rand.Intn(len(r))].(string), nil
		}
	}

	return "", errors.New("No valid response with the name " + name + " exists.")
}
