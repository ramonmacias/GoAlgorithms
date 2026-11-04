package structures

import (
	"reflect"
	"testing"
)

func TestBinaryHeapRemovesValuesInAscendingOrder(t *testing.T) {
	implementations := map[string]BinaryHeapModeImpl{
		"fast insert":  FastInsert,
		"fast remove":  FastRemove,
		"equilibrated": Equilibrated,
	}
	testCases := map[string]struct {
		input []int
		want  []int
	}{
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
		"duplicates": {
			input: []int{3, 1, 2, 1, 3, 2},
			want:  []int{1, 1, 2, 2, 3, 3},
		},
		"negative values": {
			input: []int{-2, -5, -1, -4, -3},
			want:  []int{-5, -4, -3, -2, -1},
		},
		"mixed signs and zero": {
			input: []int{7, -3, 0, 7, -10, 2},
			want:  []int{-10, -3, 0, 2, 7, 7},
		},
	}

	for implementation, mode := range implementations {
		t.Run(implementation, func(t *testing.T) {
			for name, tc := range testCases {
				t.Run(name, func(t *testing.T) {
					heap := BuildBinaryHeap(mode)
					for _, value := range tc.input {
						heap.Insert(value)
					}

					got := make([]int, len(tc.input))
					for i := range got {
						got[i] = heap.RemoveMin()
					}

					if !reflect.DeepEqual(got, tc.want) {
						t.Fatalf("RemoveMin() sequence = %v, want %v", got, tc.want)
					}
				})
			}
		})
	}
}

func TestBinaryHeapSupportsInterleavedOperations(t *testing.T) {
	for name, mode := range map[string]BinaryHeapModeImpl{
		"fast insert":  FastInsert,
		"fast remove":  FastRemove,
		"equilibrated": Equilibrated,
	} {
		t.Run(name, func(t *testing.T) {
			heap := BuildBinaryHeap(mode)
			heap.Insert(5)
			heap.Insert(1)

			if got := heap.RemoveMin(); got != 1 {
				t.Fatalf("first RemoveMin() = %d, want 1", got)
			}

			heap.Insert(3)
			heap.Insert(0)
			for i, want := range []int{0, 3, 5} {
				if got := heap.RemoveMin(); got != want {
					t.Fatalf("RemoveMin() #%d = %d, want %d", i+2, got, want)
				}
			}
		})
	}
}
