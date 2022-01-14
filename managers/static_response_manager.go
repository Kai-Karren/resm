package managers

import "errors"

type StaticResponseManager struct {
	Name_to_response map[string]string
}

func (s *StaticResponseManager) GetResponse(name string) (string, error) {

	if response, ok := s.Name_to_response[name]; ok {
		return response, nil
	}

	return "", errors.New("No response with the name " + name + " exists.")
}
