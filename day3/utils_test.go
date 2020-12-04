package main

import (
	"fmt"
	"testing"
)

func TestFromFileData(t *testing.T) {
	testcases := []struct {
		filename     string
		expectedRows int
		slopeChecks  []struct {
			slope         slope
			expectedTrees int
		}
		expecedProduct int
	}{
		{
			filename:     "test_input.txt",
			expectedRows: 11,
			slopeChecks: []struct {
				slope         slope
				expectedTrees int
			}{
				{slope{1, 1}, 2},
				{slope{3, 1}, 7},
				{slope{5, 1}, 3},
				{slope{7, 1}, 4},
				{slope{1, 2}, 2},
			},
			expecedProduct: 336,
		},
		{
			filename:     "my_input.txt",
			expectedRows: 323,
			slopeChecks: []struct {
				slope         slope
				expectedTrees int
			}{
				{slope{1, 1}, 70},
				{slope{3, 1}, 220},
				{slope{5, 1}, 63},
				{slope{7, 1}, 76},
				{slope{1, 2}, 29},
			},
			expecedProduct: 2138320800,
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("Tests with file %s", testcase.filename), func(t *testing.T) {
			area, err := readData(testcase.filename)
			if err != nil {
				t.Error(err)
			}
			// fmt.Printf("area:\n%v", area)

			rows := area.rows()
			if rows != testcase.expectedRows {
				t.Errorf("Unexpected number of rows: actual %v, expected %v", rows, testcase.expectedRows)
			}

			product := 1
			for _, slopeCheck := range testcase.slopeChecks {
				slope := slopeCheck.slope
				t.Run(fmt.Sprintf("Slope %d %d should found %d trees", slope.x, slope.y, slopeCheck.expectedTrees), func(t *testing.T) {
					trees := countTrees(area, slope)
					if trees != slopeCheck.expectedTrees {
						t.Errorf("Unexpected number of trees for slope %d %d: actual %v, expected %v", slope.x, slope.y, trees, slopeCheck.expectedTrees)
					}
					product *= trees
				})
			}

			if product != testcase.expecedProduct {
				t.Errorf("Unexpected product of trees: actual %v, expected %v", product, testcase.expecedProduct)
			}
		})
	}
}
