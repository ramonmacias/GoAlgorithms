package sorting

// Insertion sorts the slice in place by growing a sorted prefix from left to
// right. At each iteration, it takes the next value and swaps it left past
// every larger value until it reaches its correct position. This is similar to
// inserting a playing card into an already sorted hand.
//
// If the slice is already sorted, the inner loop makes one comparison per
// element and no swaps, giving a best-case time complexity of O(n). If the
// slice is reverse-sorted, element i must cross all i preceding elements. The
// total work is therefore 1 + 2 + ... + (n-1) = n(n-1)/2, giving worst-case
// O(n^2) time. The average case is also O(n^2). Because it swaps only adjacent
// out-of-order values, this implementation is stable and uses O(1) extra space.
//
// The method mirrors the much older practice of sorting cards by hand. Its
// earliest published description in the context of computer sorting is
// generally credited to John Mauchly's 1946 Moore School lectures, where he
// discussed both straight insertion and binary insertion.
func Insertion(a []int) []int {
	for i := 0; i < len(a); i++ {
		j := i
		for j > 0 && a[j] < a[j-1] {
			tmp := a[j]
			a[j] = a[j-1]
			a[j-1] = tmp

			j--
		}
	}
	return a
}
