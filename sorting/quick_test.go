package sorting_test

import (
	"math/rand"
	"slices"
	"sorting"
	"testing"
)

func TestQuick(t *testing.T) {
	testQuick(t, sorting.Quick)
}

func TestQuickImproved(t *testing.T) {
	testQuick(t, sorting.QuickImproved)
}

func testQuick(t *testing.T, sortFunc func([]int)) {
	t.Helper()
	cases := map[string][]int{
		"nil":          nil,
		"empty":        {},
		"single":       {42},
		"two reversed": {2, 1},
		"sorted":       {1, 2, 3, 4, 5},
		"reversed":     {5, 4, 3, 2, 1},
		"duplicates":   {3, 1, 2, 3, 1, 2},
		"equal":        {8, 8, 8, 8},
		"many equal":   make([]int, 10000),
		"mixed signs":  {0, -10, 7, -3, 7, 2},
	}
	rng := rand.New(rand.NewSource(1))
	random := make([]int, 1000)
	for i := range random {
		random[i] = rng.Intn(201) - 100
	}
	cases["random"] = random
	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			want := slices.Clone(input)
			slices.Sort(want)
			sortFunc(input)
			if !slices.Equal(input, want) {
				t.Fatalf("Quick = %v, want %v", input, want)
			}
		})
	}
}
