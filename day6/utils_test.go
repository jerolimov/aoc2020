package main

import (
	"testing"
)

func TestReadFromFile(t *testing.T) {
	testcases := []struct {
		filename               string
		expectedNumberOfGroups int
		expectedSum            int
	}{
		{
			filename:               "test_input.txt",
			expectedNumberOfGroups: 5,
			expectedSum:            11,
		},
		{
			filename:               "my_input.txt",
			expectedNumberOfGroups: 456,
			expectedSum:            6310,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.filename, func(t *testing.T) {
			declarations, err := readGroupDeclarationFromFile(testcase.filename)
			if err != nil {
				t.Error(err)
				return
			}

			if len(declarations) != testcase.expectedNumberOfGroups {
				t.Errorf("Unexpected number of groups: actual %v, expected %v", len(declarations), testcase.expectedNumberOfGroups)
			}

			sum := 0
			for _, declaration := range declarations {
				for _, answer := range declaration {
					if answer {
						sum++
					}
				}
			}
			if sum != testcase.expectedSum {
				t.Errorf("Unexpected sum: actual %v, expected %v", sum, testcase.expectedSum)
			}
		})
	}
}
