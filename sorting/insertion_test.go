package sorting_test

import (
	"sorting"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertion(t *testing.T) {
	testCases := map[string]struct {
		input          []int
		expectedOutput []int
	}{
		// This is O(n) complexity
		"best case": {
			input:          []int{1, 2, 3, 4},
			expectedOutput: []int{1, 2, 3, 4},
		},
		// This is O(n^2) complexity
		"worst case": {
			input:          []int{4, 3, 2, 1},
			expectedOutput: []int{1, 2, 3, 4},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			sorted := sorting.Insertion(tc.input)
			assert.Equal(t, tc.expectedOutput, sorted)
		})
	}
}
