package main

import (
	"bufio"
	"io"
	"os"
)

type declarationAnswers [26]bool

func readGroupDeclarationFromFile(filename string) ([]declarationAnswers, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return readGroupDeclaration(file)
}

func readGroupDeclaration(reader io.Reader) ([]declarationAnswers, error) {
	declarations := make([]declarationAnswers, 0)
	scanner := bufio.NewScanner(reader)

	declaration := declarationAnswers{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			declarations = append(declarations, declaration)
			declaration = declarationAnswers{}
		} else {
			for _, b := range []byte(line) {
				index := b - 'a'
				declaration[index] = true
			}
		}
	}
	declarations = append(declarations, declaration)
	if err := scanner.Err(); err != nil {
		return declarations, err
	}
	return declarations, nil
}
