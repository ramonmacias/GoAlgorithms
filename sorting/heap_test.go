package sorting_test

import (
	"slices"
	"sorting"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	testCases := map[string]struct {
		input []int
		want  []int
	}{
		"empty slice": {
			input: []int{},
			want:  []int{},
		},
		"single value": {
			input: []int{42},
			want:  []int{42},
		},
		"already sorted": {
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		"reverse sorted": {
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		"duplicate values": {
			input: []int{3, 1, 2, 3, 1, 2},
			want:  []int{1, 1, 2, 2, 3, 3},
		},
		"negative values": {
			input: []int{-2, -5, -1, -3, -4},
			want:  []int{-5, -4, -3, -2, -1},
		},
		"mixed signs and zero": {
			input: []int{7, -3, 0, 7, -10, 2},
			want:  []int{-10, -3, 0, 2, 7, 7},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			original := slices.Clone(tc.input)

			assert.Equal(t, tc.want, sorting.Heap(tc.input))
			assert.Equal(t, original, tc.input, "Heap must not modify its input")
		})
	}
}
