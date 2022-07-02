package responses

import (
	"fmt"
	"strings"
)

func FillSlots(response string, slots map[string]string) string {

	if containsVariable(response) {

		possibleKeys := make([]string, 0, len(slots))

		for key := range slots {
			possibleKeys = append(possibleKeys, key)
		}

		foundKeys := containsKeys(response, possibleKeys)

		response = replaceVariableWithSlotValue(response, foundKeys, slots)

	}

	return response

}

func containsVariable(response string) bool {

	return strings.Contains(response, "$") || (strings.Contains(response, "{") && strings.Contains(response, "}"))

}

func containsKeys(response string, keys []string) []string {

	foundKeys := make([]string, 0)

	for _, key := range keys {

		fmt.Println(key)

		if strings.Contains(response, key) {
			foundKeys = append(foundKeys, key)
		}

	}

	return foundKeys

}

func replaceVariableWithSlotValue(response string, foundVariables []string, slots map[string]string) string {

	for _, variable := range foundVariables {

		response = strings.Replace(response, "$"+variable, slots[variable], -1)
		response = strings.Replace(response, "{"+variable+"}", slots[variable], -1)

	}

	return response

}
