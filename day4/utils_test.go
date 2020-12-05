package main

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	actualPassports, err := readFile("my_input.txt")
	expectedPassports := []passport{
		{
			// ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
			// byr:1937 iyr:2017 cid:147 hgt:183cm
			"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd",
			"byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm",
		},
		{
			// iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
			// hcl:#cfa07d byr:1929
			"iyr": "2013", "ecl": "amb", "cid": "350", "eyr": "2023", "pid": "028048884",
			"hcl": "#cfa07d", "byr": "1929",
		},
		{
			// hcl:#ae17e1 iyr:2013
			// eyr:2024
			// ecl:brn pid:760753108 byr:1931
			// hgt:179cm
			"hcl": "#ae17e1", "iyr": "2013",
			"eyr": "2024",
			"ecl": "brn", "pid": "760753108", "byr": "1931",
			"hgt": "179cm",
		},
		{
			// hcl:#cfa07d eyr:2025 pid:166559648
			// iyr:2011 ecl:brn hgt:59in
			"hcl": "#cfa07d", "eyr": "2025", "pid": "166559648",
			"iyr": "2011", "ecl": "brn", "hgt": "59in",
		},
	}
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(actualPassports, expectedPassports) {
		t.Errorf("Unexpected value: actual \n%v\n, expected \n%v\n", actualPassports, expectedPassports)
	}
}

func TestDataPart1(t *testing.T) {
	testcases := []struct {
		filename                 string
		expectedPassports        int
		expectedValidPassports   int
		expectedInvalidPassports int
	}{
		{
			filename:                 "my_input.txt",
			expectedPassports:        4,
			expectedValidPassports:   2,
			expectedInvalidPassports: 2,
		},
		{
			filename:                 "test_input.txt",
			expectedPassports:        260,
			expectedValidPassports:   222,
			expectedInvalidPassports: 38,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.filename, func(t *testing.T) {
			passports, err := readFile(testcase.filename)
			if err != nil {
				t.Error(err)
			}

			valid := 0
			invalid := 0
			for _, passport := range passports {
				if isPassportValidPart1(passport) {
					valid++
				} else {
					invalid++
				}
			}

			if len(passports) != testcase.expectedPassports {
				t.Errorf("Unexpected passport count: actual %v, expected %v", len(passports), testcase.expectedPassports)
			}
			if valid != testcase.expectedValidPassports {
				t.Errorf("Unexpected valid passport count: actual %v, expected %v", valid, testcase.expectedValidPassports)
			}
			if invalid != testcase.expectedInvalidPassports {
				t.Errorf("Unexpected invalid passport count: actual %v, expected %v", invalid, testcase.expectedInvalidPassports)
			}
		})
	}
}

func TestDataPart2(t *testing.T) {
	testcases := []struct {
		filename                 string
		expectedPassports        int
		expectedValidPassports   int
		expectedInvalidPassports int
	}{
		{
			filename:                 "my_input.txt",
			expectedPassports:        4,
			expectedValidPassports:   2,
			expectedInvalidPassports: 2,
		},
		{
			filename:                 "test_input.txt",
			expectedPassports:        260,
			expectedValidPassports:   140,
			expectedInvalidPassports: 120,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.filename, func(t *testing.T) {
			passports, err := readFile(testcase.filename)
			if err != nil {
				t.Error(err)
			}

			valid := 0
			invalid := 0
			for _, passport := range passports {
				if isPassportValidPart2(passport) {
					valid++
				} else {
					invalid++
				}
			}

			if len(passports) != testcase.expectedPassports {
				t.Errorf("Unexpected passport count: actual %v, expected %v", len(passports), testcase.expectedPassports)
			}
			if valid != testcase.expectedValidPassports {
				t.Errorf("Unexpected valid passport count: actual %v, expected %v", valid, testcase.expectedValidPassports)
			}
			if invalid != testcase.expectedInvalidPassports {
				t.Errorf("Unexpected invalid passport count: actual %v, expected %v", invalid, testcase.expectedInvalidPassports)
			}
		})
	}
}
