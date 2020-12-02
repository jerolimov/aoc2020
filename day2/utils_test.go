package main

import (
	"testing"
)

func TestValidatePart1(t *testing.T) {
	assertValidatePart1(t, condition{1, 3, "a"}, "abcde", true)
	assertValidatePart1(t, condition{1, 3, "b"}, "cdefg", false)
	assertValidatePart1(t, condition{2, 9, "c"}, "ccccccccc", true)
}

func assertValidatePart1(t *testing.T, condition condition, password string, expectedValidation bool) {
	actualValidation := validatePart1(condition, password)
	if actualValidation != expectedValidation {
		t.Errorf("Unexpected validation for condition %v with password %s result: actual %v, expected %v", condition, password, actualValidation, expectedValidation)
	}
}

func TestValidatePart2(t *testing.T) {
	assertValidatePart2(t, condition{1, 3, "a"}, "abcde", true)
	assertValidatePart2(t, condition{1, 3, "b"}, "cdefg", false)
	assertValidatePart2(t, condition{2, 9, "c"}, "ccccccccc", false)
}

func assertValidatePart2(t *testing.T, condition condition, password string, expectedValidation bool) {
	actualValidation := validatePart2(condition, password)
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
		if validatePart1(entry.condition, entry.password) {
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

func TestMyInputPart1(t *testing.T) {
	entries, err := readData("my_input.txt")
	if err != nil {
		t.Error(err)
	}
	valid := 0
	invalid := 0
	for _, entry := range entries {
		if validatePart1(entry.condition, entry.password) {
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

func TestMyInputPart2(t *testing.T) {
	entries, err := readData("my_input.txt")
	if err != nil {
		t.Error(err)
	}
	valid := 0
	invalid := 0
	for _, entry := range entries {
		if validatePart2(entry.condition, entry.password) {
			valid++
		} else {
			invalid++
		}
	}
	if valid != 588 {
		t.Errorf("Unexpected number of valid entries: actual %v, expected %v", valid, 588)
	}
	if invalid != 412 {
		t.Errorf("Unexpected number of invalid entries: actual %v, expected %v", invalid, 412)
	}
}
