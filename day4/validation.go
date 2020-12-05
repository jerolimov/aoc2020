package main

import (
	"regexp"
	"strconv"
	"strings"
)

func isPassportValidPart1(passport passport) bool {
	if passport["cid"] == "" {
		return len(passport) == 7
	}
	return len(passport) == 8
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func isBirthYearValid(input string) bool {
	year, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return year >= 1920 && year <= 2002
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func isIssueYearValid(input string) bool {
	year, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return year >= 2010 && year <= 2020
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func isExpirationYearValid(input string) bool {
	year, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return year >= 2020 && year <= 2030
}

// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
func isHeightValid(input string) bool {
	if strings.HasSuffix(input, "cm") {
		year, err := strconv.Atoi(input[0 : len(input)-2])
		if err != nil {
			return false
		}
		return year >= 150 && year <= 193
	} else if strings.HasSuffix(input, "in") {
		year, err := strconv.Atoi(input[0 : len(input)-2])
		if err != nil {
			return false
		}
		return year >= 59 && year <= 76
	} else {
		return false
	}
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func isHairColorValid(input string) bool {
	return regexp.MustCompile("^#[0-9a-f]{6}$").MatchString(input)
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func isEyeColorValid(input string) bool {
	return regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$").MatchString(input)
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func isPassportIDValid(input string) bool {
	return regexp.MustCompile("^[0-9]{9}$").MatchString(input)
}

func isPassportValidPart2(passport passport) bool {
	return isBirthYearValid(passport["byr"]) &&
		isIssueYearValid(passport["iyr"]) &&
		isExpirationYearValid(passport["eyr"]) &&
		isHeightValid(passport["hgt"]) &&
		isHairColorValid(passport["hcl"]) &&
		isEyeColorValid(passport["ecl"]) &&
		isPassportIDValid(passport["pid"])
}
