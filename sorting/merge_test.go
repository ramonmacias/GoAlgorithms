package sorting_test

import (
	"sorting"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	testCases := map[string]struct {
		input          []int
		expectedOutput []int
	}{
		"best case": {
			input:          []int{1, 2, 3, 4},
			expectedOutput: []int{1, 2, 3, 4},
		},
		"worst case": {
			input:          []int{4, 3, 2, 1},
			expectedOutput: []int{1, 2, 3, 4},
		},
		"empty slice": {
			input:          []int{},
			expectedOutput: []int{},
		},
		"single value": {
			input:          []int{42},
			expectedOutput: []int{42},
		},
		"two unsorted values": {
			input:          []int{2, 1},
			expectedOutput: []int{1, 2},
		},
		"odd number of values": {
			input:          []int{5, 1, 4, 2, 3},
			expectedOutput: []int{1, 2, 3, 4, 5},
		},
		"duplicate values": {
			input:          []int{3, 1, 2, 3, 1, 2},
			expectedOutput: []int{1, 1, 2, 2, 3, 3},
		},
		"negative values": {
			input:          []int{-2, -5, -1, -3, -4},
			expectedOutput: []int{-5, -4, -3, -2, -1},
		},
		"mixed signs and zero": {
			input:          []int{0, -10, 7, -3, 7, 2},
			expectedOutput: []int{-10, -3, 0, 2, 7, 7},
		},
		"all values equal": {
			input:          []int{8, 8, 8, 8},
			expectedOutput: []int{8, 8, 8, 8},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			sorted := sorting.Merge(tc.input)
			assert.Equal(t, tc.expectedOutput, sorted)
		})
	}
}
