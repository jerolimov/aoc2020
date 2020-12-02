package main

import (
	"testing"
)

func TestValidate(t *testing.T) {
	assertValidate(t, condition{1, 3, "a"}, "abcde", true)
	assertValidate(t, condition{1, 3, "b"}, "cdefg", false)
	assertValidate(t, condition{2, 9, "c"}, "ccccccccc", true)
}

func assertValidate(t *testing.T, condition condition, password string, expectedValidation bool) {
	actualValidation := validate(condition, password)
	if actualValidation != expectedValidation {
		t.Errorf("Unexpected validation for condition %v with password %s result: actual %v, expected %v", condition, password, actualValidation, expectedValidation)
	}
}

func TestInput(t *testing.T) {
	entries, err := readData("test_input.txt")
	if err != nil {
		t.Error(err)
	}
	valid := 0
	invalid := 0
	for _, entry := range entries {
		if validate(entry.condition, entry.password) {
			valid++
		} else {
			invalid++
		}
	}
	if valid != 2 {
		t.Errorf("Unexpected number of valid entries: actual %v, expected %v", valid, 2)
	}
	if invalid != 1 {
		t.Errorf("Unexpected number of invalid entries: actual %v, expected %v", invalid, 1)
	}
}

func TestMyInput(t *testing.T) {
	entries, err := readData("my_input.txt")
	if err != nil {
		t.Error(err)
	}
	valid := 0
	invalid := 0
	for _, entry := range entries {
		if validate(entry.condition, entry.password) {
			valid++
		} else {
			invalid++
		}
	}
	if valid != 636 {
		t.Errorf("Unexpected number of valid entries: actual %v, expected %v", valid, 636)
	}
	if invalid != 364 {
		t.Errorf("Unexpected number of invalid entries: actual %v, expected %v", invalid, 364)
	}
}
