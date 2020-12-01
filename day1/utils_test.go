package main

import "testing"

func TestInput(t *testing.T) {
	numbers, err := readNumbers("test_input.txt")
	if err != nil {
		t.Error(err)
	}

	actualSum, err := findSum(numbers, 2020)
	if err != nil {
		t.Error(err)
	}
	expectedSum := 514579
	if actualSum != expectedSum {
		t.Errorf("Unexpected findSum result: actual %d, expected %d", actualSum, expectedSum)
	}

	actualSumOf3, err := findSumOf3(numbers, 2020)
	if err != nil {
		t.Error(err)
	}
	expectedSumOf3 := 241861950
	if actualSumOf3 != expectedSumOf3 {
		t.Errorf("Unexpected findSum result: actual %d, expected %d", actualSumOf3, expectedSumOf3)
	}
}
