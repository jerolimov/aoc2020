package main

import (
	"bufio"
	"os"
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
			highestSeatID := -1

			file, err := os.Open(testcase.filename)
			if err != nil {
				t.Error(err)
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
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
