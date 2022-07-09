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

func TestInMemoryAddResponse(t *testing.T) {

	storage := NewInMemoryResponseStorage()

	storage.AddResponse("test", []string{"This is a test"})

	response, err := storage.GetFirstResponse("test")

	if err != nil {
		t.Fail()
	}

	if response != "This is a test" {
		t.Fail()
	}

}
