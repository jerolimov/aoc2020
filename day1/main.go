package main

import (
	"fmt"
	"log"
)

func main() {
	numbers, err := readNumbers("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d numbers\n", len(numbers))
	sum, err := findSum(numbers, 2020)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  sum: %d\n", sum)
	sumOf3, err := findSumOf3(numbers, 2020)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  sumOf3: %d\n", sumOf3)

	numbers, err = readNumbers("my_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d numbers\n", len(numbers))
	sum, err = findSum(numbers, 2020)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  sum: %d\n", sum)
	sumOf3, err = findSumOf3(numbers, 2020)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  sumOf3: %d\n", sumOf3)
}
