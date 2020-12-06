package main

import (
	"testing"
)

func TestReadFromFile(t *testing.T) {
	testcases := []struct {
		filename               string
		expectedNumberOfGroups int
		expectedAnyoneSum      int
	}{
		{
			filename:               "test_input.txt",
			expectedNumberOfGroups: 5,
			expectedAnyoneSum:      11,
		},
		{
			filename:               "my_input.txt",
			expectedNumberOfGroups: 456,
			expectedAnyoneSum:      6310,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.filename, func(t *testing.T) {
			groupedDeclarations, err := readGroupedDeclarationFromFile(testcase.filename)
			if err != nil {
				t.Error(err)
				return
			}

			if len(groupedDeclarations) != testcase.expectedNumberOfGroups {
				t.Errorf("Unexpected number of groups: actual %v, expected %v", len(groupedDeclarations), testcase.expectedNumberOfGroups)
			}

			anyoneSum := 0
			for _, groupDeclaration := range groupedDeclarations {
				for _, answer := range reduceAnyone(groupDeclaration) {
					if answer {
						anyoneSum++
					}
				}
			}
			if anyoneSum != testcase.expectedAnyoneSum {
				t.Errorf("Unexpected anyone sum: actual %v, expected %v", anyoneSum, testcase.expectedAnyoneSum)
			}
		})
	}
}
