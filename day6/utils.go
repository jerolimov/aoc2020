package main

import (
	"bufio"
	"io"
	"os"
)

type declarationAnswers [26]bool

type groupDeclaration []declarationAnswers

func readGroupedDeclarationFromFile(filename string) ([]groupDeclaration, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return readGroupedDeclaration(file)
}

func readGroupedDeclaration(reader io.Reader) ([]groupDeclaration, error) {
	groupedDeclarations := make([]groupDeclaration, 0)

	scanner := bufio.NewScanner(reader)

	groupDeclaration := make([]declarationAnswers, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			groupedDeclarations = append(groupedDeclarations, groupDeclaration)
			groupDeclaration = make([]declarationAnswers, 0)
		} else {
			declaration := declarationAnswers{}
			for _, b := range []byte(line) {
				index := b - 'a'
				declaration[index] = true
			}
			groupDeclaration = append(groupDeclaration, declaration)
		}
	}

	groupedDeclarations = append(groupedDeclarations, groupDeclaration)

	if err := scanner.Err(); err != nil {
		return groupedDeclarations, err
	}
	return groupedDeclarations, nil
}

func reduceAnyone(declarations groupDeclaration) declarationAnswers {
	result := declarationAnswers{}
	for _, declaration := range declarations {
		for index, answer := range declaration {
			if answer {
				result[index] = true
			}
		}
	}
	return result
}
