package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readData(filename string) (area, error) {
	var area area
	area.addRow()

	file, err := os.Open(filename)
	if err != nil {
		return area, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		byte, err := reader.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			return area, err
		}

		switch byte {
		case '.':
			area.addFree()
		case '#':
			area.addTree()
		case '\n':
			area.addRow()
		default:
			return area, fmt.Errorf("Invalid byte '%b' in file %s", byte, filename)
		}
	}

	return area, nil
}

func countTrees(area area, slope slope) int {
	count := 0
	x := 0
	y := 0
	for {
		if y >= area.rows() {
			break
		}
		if area.isTree(x, y) {
			count++
		}
		x += slope.x
		y += slope.y
	}
	return count
}
