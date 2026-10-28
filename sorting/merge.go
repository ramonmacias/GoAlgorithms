package sorting

// Merge sorts a slice using divide and conquer. It repeatedly divides the
// input into two halves until each part contains at most one value, then merges
// the sorted halves by repeatedly taking their smallest remaining value.
//
// Splitting creates about log2(n) levels of recursion, and merging processes
// all n values once at every level. Its best-, average-, and worst-case time
// complexity are therefore all O(n log n); unlike insertion sort, an already
// sorted input does not reduce the work. This implementation returns a newly
// allocated sorted slice for inputs of two or more values. Its merge buffers
// use O(n) auxiliary space, while recursion uses O(log n) stack space.
//
// Equal values are taken from the left half first, preserving their original
// relative order and making the algorithm stable.
//
// Merge sort was devised by John von Neumann in 1945 while working on EDVAC,
// one of the earliest stored-program computers. His original program used the
// bottom-up form of the algorithm rather than this now-common recursive form.
func Merge(a []int) []int {
	if len(a) < 2 {
		return a
	}
	n := len(a)
	b := a[:n/2]
	c := a[n/2:]

	b = Merge(b)
	c = Merge(c)

	return merge(b, c)
}

func merge(a, b []int) []int {
	n := len(a)
	m := len(b)
	c := make([]int, n+m)

	i, j, k := 0, 0, 0

	for i < n || j < m {
		if j == m || (i < n && a[i] <= b[j]) {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
		k++
	}
	return c
}
