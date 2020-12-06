package main

import "math"

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
