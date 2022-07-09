package rasa

import (
	"testing"
)

func TestGetLastUtteranceName(t *testing.T) {

	events := []Event{
		{
			Event: "action",
			Name:  "utter_test",
		},
	}

	tracker := Tracker{
		SenderId: "42",
		Slots:    make(map[string]string),
		Events:   events,
	}

	lastUtteranceName, err := tracker.GetLastUtteranceName()

	if err != nil {
		t.Fail()
	}

	if lastUtteranceName != "utter_test" {
		t.Fail()
	}

}

func TestGetLastUtteranceNameMoreComplex(t *testing.T) {

	events := []Event{
		{
			Event: "action",
			Name:  "utter_welcome",
		},
		{
			Event: "action",
			Name:  "utter_test",
		},
		{
			Event: "action",
			Name:  "utter_bye",
		},
		{
			Event: "bot",
		},
	}

	tracker := Tracker{
		SenderId: "42",
		Slots:    make(map[string]string),
		Events:   events,
	}

	lastUtteranceName, err := tracker.GetLastUtteranceName()

	if err != nil {
		t.Fail()
	}

	if lastUtteranceName != "utter_bye" {
		t.Fail()
	}

}

func TestContainsAction_shouldReturnTrue_whenActionInEvents(t *testing.T) {

	events := []Event{
		{
			Event: "action",
			Name:  "utter_welcome",
		},
		{
			Event: "action",
			Name:  "utter_test",
		},
		{
			Event: "action",
			Name:  "utter_bye",
		},
		{
			Event: "bot",
		},
	}

	tracker := Tracker{
		SenderId: "42",
		Slots:    make(map[string]string),
		Events:   events,
	}

	result := tracker.ContainsAction("utter_test")

	if result == false {
		t.Fail()
	}

}

func TestContainsAction_shouldReturnFalse_whenActionNotInEvents(t *testing.T) {

	events := []Event{
		{
			Event: "action",
			Name:  "utter_welcome",
		},
		{
			Event: "action",
			Name:  "utter_test",
		},
		{
			Event: "action",
			Name:  "utter_bye",
		},
		{
			Event: "bot",
		},
	}

	tracker := Tracker{
		SenderId: "42",
		Slots:    make(map[string]string),
		Events:   events,
	}

	result := tracker.ContainsAction("action_lookup_user_data")

	if result == true {
		t.Fail()
	}

}
