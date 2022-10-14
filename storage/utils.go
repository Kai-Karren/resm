package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func AddResponsesFromDir(storage ResponseStorage, dir string) {

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if isJSON(file.Name()) {
			AddResponsesFromJson(storage, dir+"/"+file.Name())
		} else if file.IsDir() {
			AddResponsesFromDir(storage, dir+"/"+file.Name())
		}

	}

}

func isJSON(fileName string) bool {

	return strings.HasSuffix(fileName, ".json")

}

func AddResponsesFromJson(storage ResponseStorage, fileName string) {

	responses := ReadJsonFile(fileName)

	for key, value := range responses {
		storage.AddResponses(key, value)
	}

}

func ReadJsonFile(fileName string) map[string][]string {

	data, err := os.ReadFile(fileName)

	check(err)

	var rawResponses map[string]interface{}

	json.Unmarshal([]byte(data), &rawResponses)

	responses := convertJsonToResponseMap(rawResponses)

	return responses

}

func convertJsonToResponseMap(jsonMap map[string]interface{}) map[string][]string {

	responseMap := make(map[string][]string)

	for key, value := range jsonMap {

		switch response := value.(type) {
		case string:
			responseMap[key] = []string{response}

		case []interface{}:

			assertedType := value.([]interface{})

			responses := make([]string, len(assertedType))

			for i, v := range assertedType {
				responses[i] = v.(string)
			}

			responseMap[key] = responses

		}

	}

	return responseMap

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
