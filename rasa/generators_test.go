package rasa

import (
	"testing"
)

func TestDistributedResponseGenerator(t *testing.T) {

	firstGenerator := NewCustomResponseGenerator()

	exampleHandler := func(request NlgRequest) (NlgResponse, error) {
		return NewRasaNlgResponse("This is a custom response."), nil
	}

	firstGenerator.AddHandler("utter_first", exampleHandler)

	secondGenerator := NewCustomResponseGenerator()

	secondHandler := func(request NlgRequest) (NlgResponse, error) {

		if request.Channel.Name == "Twilio" {
			return NewRasaNlgResponse("Twilio response"), nil
		}

		return NewRasaNlgResponse("This a response from the second handler."), nil
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
		return NewRasaNlgResponse("This is a custom response."), nil
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
