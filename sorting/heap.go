package sorting

import "structures"

// Heap returns the values in ascending order using a binary min-heap. It first
// inserts every value, then repeatedly removes the current minimum. Each insert
// and removal takes O(log n), so sorting takes O(n log n) in the best, average,
// and worst cases. The heap and result slices require O(n) extra space, and the
// input slice is left unchanged.
//
// Heapsort and the array-backed binary heap were introduced by J. W. J.
// Williams in 1964 in "Algorithm 232: Heapsort." Williams's original algorithm
// rearranged the input array in place. This implementation uses the same heap
// ordering idea through a separate priority queue, favoring clarity over the
// original algorithm's O(1) auxiliary-space property.
func Heap(a []int) []int {
	binaryHeap := structures.BuildBinaryHeap(structures.Equilibrated)
	for _, val := range a {
		binaryHeap.Insert(val)
	}
	res := []int{}
	for range a {
		res = append(res, binaryHeap.RemoveMin())
	}
	return res
}
