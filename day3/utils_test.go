package main

import (
	"fmt"
	"testing"
)

func TestInput(t *testing.T) {
	area, err := readData("test_input.txt")
	if err != nil {
		t.Error(err)
	}
	// fmt.Printf("area:\n%v", area)
	rows := area.rows()
	if rows != 11 {
		t.Errorf("Unexpected number of rows: actual %v, expected %v", rows, 11)
	}

	testcases := []struct {
		slope         slope
		expectedTrees int
	}{
		{slope{1, 1}, 2},
		{slope{3, 1}, 7},
		{slope{5, 1}, 3},
		{slope{7, 1}, 4},
		{slope{1, 2}, 2},
	}

	product := 1
	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("Slope %d %d should found %d trees", testcase.slope.x, testcase.slope.y, testcase.expectedTrees), func(t *testing.T) {
			actualTrees := countTrees(area, testcase.slope)
			if actualTrees != testcase.expectedTrees {
				t.Errorf("Unexpected number of trees for slope %d %d: actual %v, expected %v", testcase.slope.x, testcase.slope.y, actualTrees, testcase.expectedTrees)
			}
			product *= actualTrees
		})
	}

	if product != 336 {
		t.Errorf("Unexpected product of trees: actual %v, expected %v", product, 336)
	}
}

func TestMyInput(t *testing.T) {
	area, err := readData("my_input.txt")
	if err != nil {
		t.Error(err)
	}
	// fmt.Printf("area:\n%v", area)
	rows := area.rows()
	if rows != 323 {
		t.Errorf("Unexpected number of rows: actual %v, expected %v", rows, 323)
	}

	testcases := []struct {
		slope         slope
		expectedTrees int
	}{
		{slope{1, 1}, 70},
		{slope{3, 1}, 220},
		{slope{5, 1}, 63},
		{slope{7, 1}, 76},
		{slope{1, 2}, 29},
	}

	product := 1
	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("Slope %d %d should found %d trees", testcase.slope.x, testcase.slope.y, testcase.expectedTrees), func(t *testing.T) {
			actualTrees := countTrees(area, testcase.slope)
			if actualTrees != testcase.expectedTrees {
				t.Errorf("Unexpected number of trees for slope %d %d: actual %v, expected %v", testcase.slope.x, testcase.slope.y, actualTrees, testcase.expectedTrees)
			}
			product *= actualTrees
		})
	}

	if product != 2138320800 {
		t.Errorf("Unexpected product of trees: actual %v, expected %v", product, 2138320800)
	}
}
