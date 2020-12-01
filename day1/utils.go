package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readNumbers(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, i)
	}
	if err := scanner.Err(); err != nil {
		return numbers, err
	}

	return numbers, nil
}

func findSum(numbers []int, expectedSum int) (int, error) {
	for a := 0; a < len(numbers); a++ {
		for b := a; b < len(numbers); b++ {
			sum := numbers[a] + numbers[b]
			if sum == expectedSum {
				// fmt.Printf("%d + %d = %d -- %d * %d = %d\n", numbers[a], numbers[b], sum, numbers[a], numbers[b], numbers[a]*numbers[b])
				return numbers[a] * numbers[b], nil
			}
		}
	}
	return 0, fmt.Errorf("Don't find a sum in all numbers which match %d", expectedSum)
}

func findSumOf3(numbers []int, expectedSum int) (int, error) {
	for a := 0; a < len(numbers); a++ {
		for b := a; b < len(numbers); b++ {
			for c := b; c < len(numbers); c++ {
				sum := numbers[a] + numbers[b] + numbers[c]
				if sum == expectedSum {
					// fmt.Printf("%d + %d + %d = %d -- %d * %d * %d = %d\n", numbers[a], numbers[b], numbers[c], sum, numbers[a], numbers[b], numbers[c], numbers[a]*numbers[b]*numbers[c])
					return numbers[a] * numbers[b] * numbers[c], nil
				}
			}
		}
	}
	return 0, fmt.Errorf("Don't find a sum in all numbers which match %d", expectedSum)
}
