package managers

import "errors"

type StaticResponseManager struct {
	NameToResponse map[string]string
}

func (s *StaticResponseManager) GetResponse(name string) (string, error) {

	if response, ok := s.NameToResponse[name]; ok {
		return response, nil
	}

	return "", errors.New("No response with the name " + name + " exists.")
}
