package sorting

import (
	"math/rand"
	"slices"
)

// Quick attempts to sort a in place using randomized two-way quicksort. The
// intended approach partitions around a pivot and recursively sorts each side.
//
// A correct randomized two-way quicksort takes O(n log n) time on average and
// O(n^2) in the worst case. Balanced partitions use O(log n) stack space;
// repeatedly unbalanced partitions can use O(n). Swapping across partitions
// makes the algorithm unstable.
//
// This implementation is incomplete: it skips the final element during
// partitioning and can recurse on an unchanged range when m == l. It therefore
// does not guarantee termination, and the usual complexity bounds above do
// not apply to this implementation. Use QuickImproved for a working version.
//
// Quicksort was introduced by C. A. R. (Tony) Hoare in the early 1960s. Its
// central idea is to partition before recursing, rather than merge sorted
// halves afterward as merge sort does.
func Quick(a []int) {
	var sortRange func(l, r int)
	sortRange = func(l, r int) {
		if r-l <= 1 {
			return
		}
		x := a[rand.Intn(r-l)+l]
		m := l
		for i := l; i < r-1; i++ {
			if a[i] < x {
				a[i], a[m] = a[m], a[i]
				m++
			}
		}
		sortRange(l, m)
		sortRange(m, r)
	}
	sortRange(0, len(a))
}

// QuickImproved sorts a in place, partitioning into values below, equal to,
// and above a random pivot. Only the below and above partitions need recursion.
// The equal block contains at least the pivot, so both recursive ranges shrink.
//
// Each partition scans its range once. Expected time is O(n log n) for distinct
// values, while an all-equal input takes O(n): one scan finishes the entire
// slice. The worst case remains O(n^2) if pivot choices repeatedly produce
// unbalanced partitions. Partitioning uses O(1) extra space, with expected
// O(log n) and worst-case O(n) recursion stack space. The sort is unstable
// because swaps can change the relative order of equal values.
//
// This variant combines Hoare's quicksort with Dijkstra's Dutch National Flag
// partitioning, which groups values into three regions in one pass. Skipping
// the equal region avoids repeatedly sorting duplicate keys.
//
// Reference: https://algs4.cs.princeton.edu/23quicksort/
func QuickImproved(a []int) {
	var sortRange func(l, r int)
	sortRange = func(l, r int) {
		if r-l <= 1 {
			return
		}
		x := a[rand.Intn(r-l)+l]
		lt, i, gt := l, l, r
		for i < gt {
			switch {
			case a[i] < x:
				a[lt], a[i] = a[i], a[lt]
				lt++
				i++
			case a[i] > x:
				gt--
				a[i], a[gt] = a[gt], a[i]
			default:
				i++
			}
		}
		sortRange(l, lt)
		sortRange(gt, r)
	}
	sortRange(0, len(a))
}

// QuickSelect returns the value at zero-based index k in sorted order, modifying
// a in place. It panics if k is outside [0, len(a)). Randomized three-way
// quickselect takes expected O(n) time and worst-case O(n^2). Recursion uses
// expected O(log n) stack space and O(n) in the worst case.
//
// C. A. R. (Tony) Hoare published quickselect in 1961 as "Algorithm 65: Find,"
// alongside his early quicksort work. Both algorithms partition around a pivot,
// but selection follows only the partition containing the requested rank rather
// than sorting both sides. This variant uses three-way partitioning to finish
// immediately when k falls in the block of values equal to the pivot.
func QuickSelect(a []int, k int) int {
	if k < 0 || k >= len(a) {
		panic("sorting.QuickSelect: k out of range")
	}
	var find func(l, r int) int
	find = func(l, r int) int {
		if r-l == 1 {
			return a[k]
		}
		x := a[rand.Intn(r-l)+l]
		lt, i, gt := l, l, r
		for i < gt {
			switch {
			case a[i] < x:
				a[lt], a[i] = a[i], a[lt]
				lt++
				i++
			case a[i] > x:
				gt--
				a[i], a[gt] = a[gt], a[i]
			default:
				i++
			}
		}
		switch {
		case k < lt:
			return find(l, lt)
		case k >= gt:
			return find(gt, r)
		default:
			return x
		}
	}
	return find(0, len(a))
}

// MedianOfMediansSelect returns the value at zero-based index k in sorted order,
// modifying a in place. It divides values into groups of at most five, takes
// each group's median, and recursively selects the median of those medians as a
// deterministic pivot. That pivot guarantees enough progress on every
// partition, giving O(n) time in the best, average, and worst cases. This
// implementation uses O(n) auxiliary space and O(log n) recursion depth. It
// panics when k is outside [0, len(a)).
//
// Manuel Blum, Robert W. Floyd, Vaughan Pratt, Ronald L. Rivest, and Robert E.
// Tarjan introduced this selection method as PICK in their 1973 paper "Time
// Bounds for Selection." It is now commonly called median of medians or BFPRT,
// after the authors' surnames. Their work proved for the first time that an
// arbitrary order statistic can be selected in deterministic worst-case linear
// time, unlike randomized quickselect's O(n^2) worst case.
func MedianOfMediansSelect(a []int, k int) int {
	if k < 0 || k >= len(a) {
		panic("sorting.MedianOfMediansSelect: k out of range")
	}
	var find func(l, r int) int
	find = func(l, r int) int {
		if r-l == 1 {
			return a[k]
		}
		x := medianOfMediansPivot(a[l:r])
		lt, i, gt := l, l, r
		for i < gt {
			switch {
			case a[i] < x:
				a[lt], a[i] = a[i], a[lt]
				lt++
				i++
			case a[i] > x:
				gt--
				a[i], a[gt] = a[gt], a[i]
			default:
				i++
			}
		}
		switch {
		case k < lt:
			return find(l, lt)
		case k >= gt:
			return find(gt, r)
		default:
			return x
		}
	}
	return find(0, len(a))
}

func medianOfMediansPivot(a []int) int {
	medians := make([]int, 0, (len(a)+4)/5)
	for start := 0; start < len(a); start += 5 {
		end := min(start+5, len(a))
		group := slices.Clone(a[start:end])
		Insertion(group)
		medians = append(medians, group[len(group)/2])
	}
	if len(medians) == 1 {
		return medians[0]
	}
	return MedianOfMediansSelect(medians, len(medians)/2)
}
