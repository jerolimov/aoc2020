package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type passport map[string]string

func readFile(filename string) ([]passport, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return read(file)
}

func read(reader io.Reader) ([]passport, error) {
	passports := make([]passport, 0)
	passport := make(passport)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Printf("text: %s\n", line)
		if line == "" {
			if len(passport) > 0 {
				// fmt.Printf("passport: %v\n", passport)
				passports = append(passports, passport)
			}

			passport = make(map[string]string)
		} else {
			for _, dataPair := range strings.Split(line, " ") {
				colonIndex := strings.Index(dataPair, ":")
				key := dataPair[0:colonIndex]
				value := dataPair[colonIndex+1:]
				passport[key] = value
			}
		}
	}

	if len(passport) > 0 {
		// fmt.Printf("passport: %v\n", passport)
		passports = append(passports, passport)
	}

	if err := scanner.Err(); err != nil {
		return passports, err
	}
	return passports, nil
}

func isPassportValid(passport passport) bool {
	if passport["cid"] == "" {
		return len(passport) == 7
	}
	return len(passport) == 8
}
