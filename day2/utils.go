package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type condition struct {
	atLeast int
	atMost  int
	char    string
}

type testentry struct {
	condition condition
	password  string
}

func validatePart1(condition condition, password string) bool {
	c := strings.Count(password, condition.char)
	return c >= condition.atLeast && c <= condition.atMost
}

func validatePart2(condition condition, password string) bool {
	char1 := password[condition.atLeast-1 : condition.atLeast]
	char2 := password[condition.atMost-1 : condition.atMost]
	return (char1 == condition.char) != (char2 == condition.char)
}

func readData(filename string) ([]testentry, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []testentry

	reader := bufio.NewReader(file)
	for {
		part1, _ := reader.ReadString('-')
		atLeast, _ := strconv.Atoi(part1[:len(part1)-1])
		// fmt.Printf("part1: \"%s\", atLeast: %v\n", part1, atLeast)

		part2, _ := reader.ReadString(' ')
		atMost, _ := strconv.Atoi(part2[:len(part2)-1])
		// fmt.Printf("part2: \"%s\", atMost: %v\n", part2, atMost)

		part3, _ := reader.ReadString(':')
		char := part3[0 : len(part3)-1]
		// fmt.Printf("part3: \"%s\", char: %v\n", part3, char)

		part4, err := reader.ReadString('\n')
		password := part4[1 : len(part4)-1]
		// fmt.Printf("part4: \"%s\", password: \"%s\"\n", part4[0:len(part4)-1], password)

		entries = append(entries, testentry{condition{atLeast, atMost, char}, password})

		if err == io.EOF {
			break
		} else if err != nil {
			return entries, err
		}
	}

	return entries, nil
}
