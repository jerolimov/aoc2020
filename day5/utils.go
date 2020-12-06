package main

import (
	"bufio"
	"io"
	"math"
	"os"
)

func readLinesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return readLines(file)
}

func readLines(reader io.Reader) ([]string, error) {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}

func getRow(binarySpacePartitioning string) int {
	row := 0
	pot := 0
	for i := len(binarySpacePartitioning) - 1; i >= 0; i-- {
		b := binarySpacePartitioning[i]
		if b == 'B' {
			row += int(math.Exp2(float64(pot)))
		}
		if b == 'F' || b == 'B' {
			pot++
		}
	}
	return row
}

func getColumn(binarySpacePartitioning string) int {
	column := 0
	pot := 0
	for i := len(binarySpacePartitioning) - 1; i >= 0; i-- {
		b := binarySpacePartitioning[i]
		if b == 'R' {
			column += int(math.Exp2(float64(pot)))
		}
		if b == 'L' || b == 'R' {
			pot++
		}
	}
	return column
}

func getSeatID(binarySpacePartitioning string) int {
	return getRow(binarySpacePartitioning)*8 + getColumn(binarySpacePartitioning)
}

func getFreeSeatIDs(seatIDs map[int]bool) []int {
	freeSeatIDs := []int{}
	for seatID := 0; seatID < 128*8; seatID++ {
		// free used free
		if seatIDs[seatID-1] && !seatIDs[seatID] && seatIDs[seatID+1] {
			freeSeatIDs = append(freeSeatIDs, seatID)
		}
	}
	return freeSeatIDs
}
