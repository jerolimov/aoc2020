package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := readNumbers("test_input.txt")
	fmt.Printf("Read %d numbers\n", len(numbers))
	findSum(numbers, 2020)
	fmt.Printf("\n")

	numbers = readNumbers("my_input.txt")
	fmt.Printf("Read %d numbers\n", len(numbers))
	findSum(numbers, 2020)
	fmt.Printf("\n")
}

func readNumbers(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	return numbers
}

func findSum(numbers []int, expectedSum int) {
	for a := 0; a < len(numbers); a++ {
		for b := a; b < len(numbers); b++ {
			sum := numbers[a] + numbers[b]
			if sum == expectedSum {
				fmt.Printf("%d + %d = %d -- %d * %d = %d\n", numbers[a], numbers[b], sum, numbers[a], numbers[b], numbers[a]*numbers[b])
			}
		}
	}
}
