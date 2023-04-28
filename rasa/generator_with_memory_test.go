package rasa

import (
	"testing"

	"github.com/Kai-Karren/resm/storage"
)

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

	test := []string{"This is a test", "This is another test."}
	greet := []string{"Hello", "Hey", "Hi"}

	storage.SetResponses("test", test)
	storage.SetResponses("greet", greet)

	generator := NewStaticResponseGeneratorWithMemory(
		&storage,
	)

	response, err := generator.Generate(NlgRequest{
		Response: "test",
	})

	if err != nil {
		t.Fail()
	}

	if !contains(test, response.Text) {
		t.FailNow()
	}

	response2, err := generator.Generate(NlgRequest{
		Response: "test",
	})

	if err != nil {
		t.Fail()
	}

	if !contains(test, response2.Text) {
		t.FailNow()
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

	if !contains(greet, response3.Text) {
		t.FailNow()
	}

	response4, err := generator.Generate(NlgRequest{
		Response: "greet",
	})

	if err != nil {
		t.Fail()
	}

	if !contains(greet, response4.Text) {
		t.FailNow()
	}

	if response3 == response4 {
		t.Fail()
	}

}

func TestStaticResponseGeneratorWithMemory_injectSlotsIntoResponses(t *testing.T) {

	storage := storage.NewInMemoryResponseStorage()

	expected_utter_user_names := []string{"The name is John Doe", "The user name is John Doe"}
	expected_utter_number := []string{"The number is 42", "42 is the number"}

	storage.SetResponses("utter_user_name", []string{"The name is $user_name", "The user name is {user_name}"})
	storage.SetResponses("utter_number", []string{"The number is $number", "{number} is the number"})

	generator := NewStaticResponseGeneratorWithMemory(
		&storage,
	)

	response, err := generator.Generate(NlgRequest{
		Response: "utter_user_name",
		Tracker: Tracker{
			Slots: map[string]string{
				"user_name": "John Doe",
				"number":    "42",
			},
		},
	})

	if err != nil {
		t.Fail()
	}

	if !contains(expected_utter_user_names, response.Text) {
		t.FailNow()
	}

	response2, err := generator.Generate(NlgRequest{
		Response: "utter_user_name",
		Tracker: Tracker{
			Slots: map[string]string{
				"user_name": "John Doe",
				"number":    "42",
			},
		},
	})

	if err != nil {
		t.FailNow()
	}

	if !contains(expected_utter_user_names, response2.Text) {
		t.FailNow()
	}

	if response == response2 {
		t.Fail()
	}

	response3, err := generator.Generate(NlgRequest{
		Response: "utter_number",
		Tracker: Tracker{
			Slots: map[string]string{
				"user_name": "John Doe",
				"number":    "42",
			},
		},
	})

	if err != nil {
		t.Fail()
	}

	if !contains(expected_utter_number, response3.Text) {
		t.FailNow()
	}

	response4, err := generator.Generate(NlgRequest{
		Response: "utter_number",
		Tracker: Tracker{
			Slots: map[string]string{
				"user_name": "John Doe",
				"number":    "42",
			},
		},
	})

	if err != nil {
		t.Fail()
	}

	if !contains(expected_utter_number, response4.Text) {
		t.FailNow()
	}

	if response3 == response4 {
		t.Fail()
	}

}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
