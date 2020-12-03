package main

import (
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

	slope := slope{
		x: 3,
		y: 1,
	}
	count := countTrees(area, slope)
	if count != 7 {
		t.Errorf("Unexpected number of trees: actual %v, expected %v", count, 7)
	}
}
