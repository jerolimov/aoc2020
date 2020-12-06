package main

import (
	"reflect"
	"testing"
)

func TestBinarySpacePartitioning(t *testing.T) {
	testcases := []struct {
		binarySpacePartitioning string
		expectedRow             int
		expectedColumn          int
		expectedSeatID          int
	}{
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	}

	for _, testcase := range testcases {
		t.Run(testcase.binarySpacePartitioning, func(t *testing.T) {
			if getRow(testcase.binarySpacePartitioning) != testcase.expectedRow {
				t.Errorf("Unexpected row for %s: actual %v, expected %v", testcase.binarySpacePartitioning, getRow(testcase.binarySpacePartitioning), testcase.expectedRow)
			}
			if getColumn(testcase.binarySpacePartitioning) != testcase.expectedColumn {
				t.Errorf("Unexpected column for %s: actual %v, expected %v", testcase.binarySpacePartitioning, getColumn(testcase.binarySpacePartitioning), testcase.expectedColumn)
			}
			if getSeatID(testcase.binarySpacePartitioning) != testcase.expectedSeatID {
				t.Errorf("Unexpected seat ID for %s: actual %v, expected %v", testcase.binarySpacePartitioning, getSeatID(testcase.binarySpacePartitioning), testcase.expectedSeatID)
			}
		})
	}
}

func TestBinarySpacePartitioningFromFile(t *testing.T) {
	testcases := []struct {
		filename              string
		expectedHighestSeatID int
	}{
		{
			filename:              "test_input.txt",
			expectedHighestSeatID: 820,
		},
		{
			filename:              "my_input.txt",
			expectedHighestSeatID: 890,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.filename, func(t *testing.T) {
			lines, err := readLinesFromFile(testcase.filename)
			if err != nil {
				t.Error(err)
				return
			}

			highestSeatID := -1
			for _, line := range lines {
				seatID := getSeatID(line)
				if highestSeatID < seatID {
					highestSeatID = seatID
				}
			}
			if highestSeatID != testcase.expectedHighestSeatID {
				t.Errorf("Unexpected highest seat id: actual %v, expected %v", highestSeatID, testcase.expectedHighestSeatID)
			}
		})
	}
}

func TestFreeSeatIDs(t *testing.T) {
	lines, err := readLinesFromFile("my_input.txt")
	if err != nil {
		t.Error(err)
		return
	}

	seatIDs := make(map[int]bool)
	for _, line := range lines {
		seatIDs[getSeatID(line)] = true
	}
	freeSeatIDs := getFreeSeatIDs(seatIDs)
	expectedFreeSeatIDs := []int{651}

	if !reflect.DeepEqual(freeSeatIDs, expectedFreeSeatIDs) {
		t.Errorf("Unexpected free seat IDs: actual %v, expected %v", freeSeatIDs, expectedFreeSeatIDs)
	}
}
