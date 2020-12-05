package main

import (
	"strings"
	"testing"
)

func TestIsValidPassportPart1(t *testing.T) {
	testcases := []struct {
		name          string
		passport      passport
		expectedValid bool
	}{
		{
			name: "all eight fields are present",
			passport: passport{
				// ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
				// byr:1937 iyr:2017 cid:147 hgt:183cm
				"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd",
				"byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm",
			},
			expectedValid: true,
		},
		{
			name: "missing hgt",
			passport: passport{
				// iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
				// hcl:#cfa07d byr:1929
				"iyr": "2013", "ecl": "amb", "cid": "350", "eyr": "2023", "pid": "028048884",
				"hcl": "#cfa07d", "byr": "1929",
			},
			expectedValid: false,
		},
		{
			name: "only missing field is cid",
			passport: passport{
				// hcl:#ae17e1 iyr:2013
				// eyr:2024
				// ecl:brn pid:760753108 byr:1931
				// hgt:179cm
				"hcl": "#ae17e1", "iyr": "2013",
				"eyr": "2024",
				"ecl": "brn", "pid": "760753108", "byr": "1931",
				"hgt": "179cm",
			},
			expectedValid: true,
		},
		{
			name: "missing two fields, cid and byr",
			passport: passport{
				// hcl:#cfa07d eyr:2025 pid:166559648
				// iyr:2011 ecl:brn hgt:59in
				"hcl": "#cfa07d", "eyr": "2025", "pid": "166559648",
				"iyr": "2011", "ecl": "brn", "hgt": "59in",
			},
			expectedValid: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			if isPassportValidPart1(testcase.passport) != testcase.expectedValid {
				t.Errorf("Unexpected valid status for %s: actual \n%v\n, expected \n%v\n", testcase.name, testcase.passport, testcase.expectedValid)
			}
		})
	}
}

func TestIsBirthYearValid(t *testing.T) {
	testcases := []struct {
		input         string
		expectedValid bool
	}{
		{"", false},
		{"abc", false},
		{"123", false},
		{"1980", true},
		{"2002", true},
		{"2003", false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.input, func(t *testing.T) {
			actualValid := isBirthYearValid(testcase.input)
			if actualValid != testcase.expectedValid {
				t.Errorf("Unexpected validation for %s: actual %v, expected %v", testcase.input, actualValid, testcase.expectedValid)
			}
		})
	}
}

func TestIsIssueYearValid(t *testing.T) {
	testcases := []struct {
		input         string
		expectedValid bool
	}{
		{"", false},
		{"abc", false},
		{"123", false},
		{"1980", false},
		{"2020", true},
		{"2021", false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.input, func(t *testing.T) {
			actualValid := isIssueYearValid(testcase.input)
			if actualValid != testcase.expectedValid {
				t.Errorf("Unexpected validation for %s: actual %v, expected %v", testcase.input, actualValid, testcase.expectedValid)
			}
		})
	}
}

func TestIsExpirationYearValid(t *testing.T) {
	testcases := []struct {
		input         string
		expectedValid bool
	}{
		{"", false},
		{"abc", false},
		{"123", false},
		{"1980", false},
		{"2020", true},
		{"2030", true},
		{"2031", false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.input, func(t *testing.T) {
			actualValid := isExpirationYearValid(testcase.input)
			if actualValid != testcase.expectedValid {
				t.Errorf("Unexpected validation for %s: actual %v, expected %v", testcase.input, actualValid, testcase.expectedValid)
			}
		})
	}
}

func TestIsHeightValid(t *testing.T) {
	testcases := []struct {
		input         string
		expectedValid bool
	}{
		{"", false},
		{"abc", false},
		{"123", false},
		{"60cm", false},
		{"150cm", true},
		{"200cm", false},
		{"50in", false},
		{"60in", true},
		{"180in", false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.input, func(t *testing.T) {
			actualValid := isHeightValid(testcase.input)
			if actualValid != testcase.expectedValid {
				t.Errorf("Unexpected validation for %s: actual %v, expected %v", testcase.input, actualValid, testcase.expectedValid)
			}
		})
	}
}

func TestIsHairColorValid(t *testing.T) {
	testcases := []struct {
		input         string
		expectedValid bool
	}{
		{"", false},
		{"abc", false},
		{"123123", false},
		{"#123123", true},
		{"#abcabc", true},
		{"#cdefgh", false},
		{"#ffffffff", false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.input, func(t *testing.T) {
			actualValid := isHairColorValid(testcase.input)
			if actualValid != testcase.expectedValid {
				t.Errorf("Unexpected validation for %s: actual %v, expected %v", testcase.input, actualValid, testcase.expectedValid)
			}
		})
	}
}

func TestIsEyeColorValid(t *testing.T) {
	testcases := []struct {
		input         string
		expectedValid bool
	}{
		{"", false},
		{"blu", true},
		{"blue", false},
		{"oth", true},
		{"other", false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.input, func(t *testing.T) {
			actualValid := isEyeColorValid(testcase.input)
			if actualValid != testcase.expectedValid {
				t.Errorf("Unexpected validation for %s: actual %v, expected %v", testcase.input, actualValid, testcase.expectedValid)
			}
		})
	}
}

func TestIsPassportIDValid(t *testing.T) {
	testcases := []struct {
		input         string
		expectedValid bool
	}{
		{"", false},
		{"abc", false},
		{"123", false},
		{"123456789", true},
		{"1234567890", false},
		{"asdfasdf1", false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.input, func(t *testing.T) {
			actualValid := isPassportIDValid(testcase.input)
			if actualValid != testcase.expectedValid {
				t.Errorf("Unexpected validation for %s: actual %v, expected %v", testcase.input, actualValid, testcase.expectedValid)
			}
		})
	}
}

func TestIsInvalidPassportsPart2(t *testing.T) {
	input := `
	eyr:1972 cid:100
	hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

	iyr:2019
	hcl:#602927 eyr:1967 hgt:170cm
	ecl:grn pid:012533040 byr:1946

	hcl:dab227 iyr:2012
	ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

	hgt:59cm ecl:zzz
	eyr:2038 hcl:74454a iyr:2023
	pid:3556412378 byr:2007
	`

	passports, err := read(strings.NewReader(input))
	if err != nil {
		t.Error(err)
	}

	for _, passport := range passports {
		actualValid := isPassportValidPart2(passport)
		expectedValid := false
		if actualValid != expectedValid {
			t.Errorf("Unexpected validation for passport %v: actual %v, expected %v", passport, actualValid, expectedValid)
		}
	}
}

func TestIsValidPassportsPart2(t *testing.T) {
	input := `
	pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
	hcl:#623a2f

	eyr:2029 ecl:blu cid:129 byr:1989
	iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

	hcl:#888785
	hgt:164cm byr:2001 iyr:2015 cid:88
	pid:545766238 ecl:hzl
	eyr:2022

	iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
	`

	passports, err := read(strings.NewReader(input))
	if err != nil {
		t.Error(err)
	}

	for _, passport := range passports {
		actualValid := isPassportValidPart2(passport)
		expectedValid := true
		if actualValid != expectedValid {
			t.Errorf("Unexpected validation for passport %v: actual %v, expected %v", passport, actualValid, expectedValid)
		}
	}
}
