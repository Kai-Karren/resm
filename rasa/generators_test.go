package rasa

import (
	"testing"

	"github.com/Kai-Karren/resm/storage"
)

func TestDistributedResponseGenerator(t *testing.T) {

	firstGenerator := NewCustomResponseGenerator()

	exampleHandler := func(request NlgRequest) (NlgResponse, error) {
		return NewNlgResponse("This is a custom response."), nil
	}

	firstGenerator.AddHandler("utter_first", exampleHandler)

	secondGenerator := NewCustomResponseGenerator()

	secondHandler := func(request NlgRequest) (NlgResponse, error) {

		if request.Channel.Name == "Twilio" {
			return NewNlgResponse("Twilio response"), nil
		}

		return NewNlgResponse("This a response from the second handler."), nil
	}

	secondGenerator.AddHandler("utter_second", secondHandler)

	distributed := NewDistributedResponseGenerator()

	distributed.AddGenerator(&firstGenerator)
	distributed.AddGenerator(&secondGenerator)

	response, err := distributed.Generate(NlgRequest{
		Response: "utter_first",
	})

	if err != nil {
		t.Fail()
	}

	if response.Text != "This is a custom response." {
		t.Fail()
	}

	response, err = distributed.Generate(NlgRequest{
		Response: "utter_second",
	})

	if err != nil {
		t.Fail()
	}

	if response.Text != "This a response from the second handler." {
		t.Fail()
	}

}

func TestCustomResponseGenerator_simpleCustomResponseHandler(t *testing.T) {

	generator := NewCustomResponseGenerator()

	exampleHandler := func(request NlgRequest) (NlgResponse, error) {
		return NewNlgResponse("This is a custom response."), nil
	}

	generator.AddHandler("test", exampleHandler)

	response, err := generator.Generate(NlgRequest{
		Response: "test",
	})

	if err != nil {
		t.Fail()
	}

	if response.Text != "This is a custom response." {
		t.Fail()
	}

}

func TestStaticResponseGeneratorWithMemory_singleResponseGeneration(t *testing.T) {

	storage := storage.NewInMemoryResponseStorage()

	storage.SetResponses("test", []string{"This is a test", "This is another test."})

	generator := NewStaticResponseGeneratorWithMemory(
		&storage,
	)

	response, err := generator.Generate(NlgRequest{
		Response: "test",
	})

	if err != nil {
		t.Fail()
	}

	if (response.Text != "This is a test.") != (response.Text != "This is another test") {
		t.Fail()
	}

}

func TestStaticResponseGeneratorWithMemory_generateMultipleResponses(t *testing.T) {

	storage := storage.NewInMemoryResponseStorage()

	storage.SetResponses("test", []string{"This is a test", "This is another test."})

	generator := NewStaticResponseGeneratorWithMemory(
		&storage,
	)

	response, err := generator.Generate(NlgRequest{
		Response: "test",
	})

	if err != nil {
		t.Fail()
	}

	if (response.Text != "This is a test.") != (response.Text != "This is another test") {
		t.Fail()
	}

	response2, err := generator.Generate(NlgRequest{
		Response: "test",
	})

	if err != nil {
		t.Fail()
	}

	if (response2.Text != "This is a test.") != (response2.Text != "This is another test") {
		t.Fail()
	}

	if response == response2 {
		t.Fail()
	}

}

func TestStaticResponseGeneratorWithMemory_generateMultipleResponsesForDifferentResponseNames(t *testing.T) {

	storage := storage.NewInMemoryResponseStorage()

	storage.SetResponses("test", []string{"This is a test", "This is another test."})
	storage.SetResponses("greet", []string{"Hello", "Hey", "Hi"})

	generator := NewStaticResponseGeneratorWithMemory(
		&storage,
	)

	response, err := generator.Generate(NlgRequest{
		Response: "test",
	})

	if err != nil {
		t.Fail()
	}

	if (response.Text != "This is a test.") != (response.Text != "This is another test") {
		t.Fail()
	}

	response2, err := generator.Generate(NlgRequest{
		Response: "test",
	})

	if err != nil {
		t.Fail()
	}

	if (response2.Text != "This is a test.") != (response2.Text != "This is another test") {
		t.Fail()
	}

	if response == response2 {
		t.Fail()
	}

	response3, err := generator.Generate(NlgRequest{
		Response: "greet",
	})

	if err != nil {
		t.Fail()
	}

	if (response3.Text != "Hi") != (response3.Text != "Hey") != (response3.Text != "Hello") {
		t.Fail()
	}

	response4, err := generator.Generate(NlgRequest{
		Response: "greet",
	})

	if err != nil {
		t.Fail()
	}

	if (response4.Text != "Hi") != (response4.Text != "Hey") != (response4.Text != "Hello") {
		t.Fail()
	}

	if response3 == response4 {
		t.Fail()
	}

}
