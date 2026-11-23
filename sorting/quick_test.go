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

func TestQuickSelect(t *testing.T) {
	testQuickSelect(t, sorting.QuickSelect)
}

func TestMedianOfMediansSelect(t *testing.T) {
	testQuickSelect(t, sorting.MedianOfMediansSelect)
}

func testQuickSelect(t *testing.T, selectFunc func([]int, int) int) {
	t.Helper()
	cases := map[string][]int{
		"single":       {42},
		"two reversed": {2, 1},
		"sorted":       {1, 2, 3, 4, 5},
		"reversed":     {5, 4, 3, 2, 1},
		"duplicates":   {3, 1, 2, 3, 1, 2},
		"equal":        {8, 8, 8, 8},
		"mixed signs":  {0, -10, 7, -3, 7, 2},
	}
	rng := rand.New(rand.NewSource(1))
	values := make([]int, 100)
	for i := range values {
		values[i] = rng.Intn(101) - 50
	}
	cases["random"] = values
	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			want := slices.Clone(input)
			slices.Sort(want)
			for k, expected := range want {
				if got := selectFunc(slices.Clone(input), k); got != expected {
					t.Fatalf("select(%v, %d) = %d, want %d", input, k, got, expected)
				}
			}
		})
	}
}

func TestQuickSelectInvalidIndex(t *testing.T) {
	for implementation, selectFunc := range map[string]func([]int, int) int{
		"randomized":        sorting.QuickSelect,
		"median of medians": sorting.MedianOfMediansSelect,
	} {
		t.Run(implementation, func(t *testing.T) {
			testQuickSelectInvalidIndex(t, selectFunc)
		})
	}
}

func testQuickSelectInvalidIndex(t *testing.T, selectFunc func([]int, int) int) {
	t.Helper()
	for name, tc := range map[string]struct {
		input []int
		k     int
	}{
		"nil":      {nil, 0},
		"empty":    {[]int{}, 0},
		"negative": {[]int{1}, -1},
		"past end": {[]int{1}, 1},
	} {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("expected panic for invalid k")
				}
			}()
			selectFunc(tc.input, tc.k)
		})
	}
}
