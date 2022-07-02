package responses

import "testing"

func TestFillSlotsFilleNumberSlot(t *testing.T) {

	response := "The number is $number."

	slots := make(map[string]string)

	slots["number"] = "42"

	response = FillSlots(response, slots)

	if response != "The number is 42." {
		t.Fail()
	}

}

func TestFillSlotsVerifyName(t *testing.T) {

	response := "I understood that your name is {name}, right?"

	slots := make(map[string]string)

	slots["name"] = "Kai"

	response = FillSlots(response, slots)

	if response != "I understood that your name is Kai, right?" {
		t.Fail()
	}

}

func TestFillSlotsNoSlotForVariable(t *testing.T) {

	response := "I understood that your name is {name}, right?"

	slots := make(map[string]string)

	response = FillSlots(response, slots)

	if response != "I understood that your name is {name}, right?" {
		t.Fail()
	}

}

func TestContainsVariable(t *testing.T) {

	response := "The number is $number."

	containsVariable := containsVariable(response)

	if !containsVariable {
		t.Fail()
	}

}

func TestContainsVariableNoVariable(t *testing.T) {

	response := "The number is 42."

	containsVariable := containsVariable(response)

	if containsVariable {
		t.Fail()
	}

}

func TestContainsKeys(t *testing.T) {

	response := "The number is $number."

	keys := []string{"number"}

	foundKeys := containsKeys(response, keys)

	if len(foundKeys) != 1 {
		t.Fail()
	}

	if foundKeys[0] != "number" {
		t.Fail()
	}

}

func TestReplaceVariableWithSlotValue(t *testing.T) {

	response := "The number is $number."

	foundVariables := []string{"number"}

	slots := make(map[string]string)

	slots["number"] = "42"

	response = replaceVariableWithSlotValue(response, foundVariables, slots)

	if response != "The number is 42." {
		t.Fail()
	}

}

func TestReplaceVariableWithSlotValueAlternativeFormat(t *testing.T) {

	response := "The number is {number}."

	foundVariables := []string{"number"}

	slots := make(map[string]string)

	slots["number"] = "42"

	response = replaceVariableWithSlotValue(response, foundVariables, slots)

	if response != "The number is 42." {
		t.Fail()
	}

}
