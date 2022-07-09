package storage

import "testing"

func TestInMemoryGetFirstResponse(t *testing.T) {

	responses := make(map[string][]string)
	responses["test"] = []string{"This is a test"}

	storage := InMemoryResponseStorage{
		responses: responses,
	}

	response, err := storage.GetFirstResponse("test")

	if err != nil {
		t.Fail()
	}

	if response != "This is a test" {
		t.Fail()
	}

}

func TestInMemorySetResponses(t *testing.T) {

	storage := NewInMemoryResponseStorage()

	storage.SetResponses("test", []string{"This is a test"})

	response, err := storage.GetFirstResponse("test")

	if err != nil {
		t.Fail()
	}

	if response != "This is a test" {
		t.Fail()
	}

}

func TestInMemorySetResponses_shouldOverride_whenResponseVariationAlreadyExists(t *testing.T) {

	storage := NewInMemoryResponseStorage()

	storage.SetResponses("test", []string{"This is a test"})
	storage.SetResponses("test", []string{"override successful"})

	response, err := storage.GetFirstResponse("test")

	if err != nil {
		t.Fail()
	}

	if response != "override successful" {
		t.Fail()
	}

}

func TestInMemoryAddResponses(t *testing.T) {

	storage := NewInMemoryResponseStorage()

	storage.AddResponses("test", []string{"This is a test"})

	response, err := storage.GetFirstResponse("test")

	if err != nil {
		t.Fail()
	}

	if response != "This is a test" {
		t.Fail()
	}

}

func TestInMemoryAddResponses_shouldAppendResponseVariations(t *testing.T) {

	storage := NewInMemoryResponseStorage()

	storage.AddResponses("test", []string{"This is a test"})
	storage.AddResponses("test", []string{"This is another test"})

	responses, err := storage.GetResponses("test")

	if err != nil {
		t.Fail()
	}

	if len(responses) != 2 {
		t.Fail()
	}

}

func TestInMemoryGetResponses(t *testing.T) {

	storage := NewInMemoryResponseStorage()

	storage.AddResponses("test", []string{"This is a test.", "This is a test variation."})

	responses, err := storage.GetResponses("test")

	if err != nil {
		t.Fail()
	}

	if len(responses) != 2 {
		t.Fail()
	}

}

func TestInMemoryGetResponses_shouldReturnError_whenNoResponseExists(t *testing.T) {

	storage := NewInMemoryResponseStorage()

	storage.AddResponses("test", []string{"This is a test.", "This is a test variation."})

	responses, err := storage.GetResponses("utter_welcome")

	if err == nil {
		t.Fail()
	}

	if len(responses) != 0 {
		t.Fail()
	}

}

func TestInMemoryDeleteResponse(t *testing.T) {

	storage := NewInMemoryResponseStorage()

	storage.AddResponses("test", []string{"This is a test."})

	responses, err := storage.GetResponses("test")

	if err != nil {
		t.Fail()
	}

	if len(responses) != 1 {
		t.Fail()
	}

	storage.DeleteResponses("test")

	_, err = storage.GetResponses("test")

	if err == nil {
		t.Fail()
	}

}
