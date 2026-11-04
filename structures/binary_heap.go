package structures

// BinaryHeap describes the two essential operations of a min-priority queue:
// adding a value and removing the smallest value currently stored.
//
// The binary heap was introduced by J. W. J. Williams in his 1964 paper
// "Algorithm 232: Heapsort." Williams represented a complete binary tree in an
// array, avoiding pointers, and described operations for inserting a value and
// extracting the minimum. A conventional binary heap balances both operations
// at O(log n). This file includes that balanced implementation alongside two
// simpler extremes that make either insertion or removal fast at the other's
// expense.
type BinaryHeap interface {
	Insert(x int)
	RemoveMin() int
}

// BinaryHeapModeImpl selects the time-complexity tradeoff used by a queue.
type BinaryHeapModeImpl string

const (
	// FastInsert uses an unordered slice: amortized O(1) insert and O(n) remove.
	FastInsert BinaryHeapModeImpl = "fastInsert"
	// FastRemove uses a sorted slice: O(n) insert and O(1) remove.
	FastRemove BinaryHeapModeImpl = "fastRemove"
	// Equilibrated uses a binary min-heap: O(log n) insert and remove.
	Equilibrated BinaryHeapModeImpl = "equilibrated"
)

// BuildBinaryHeap returns a slice-backed min-priority queue with the requested
// performance tradeoff. It returns nil for an unknown mode.
//
// RemoveMin must only be called when the queue contains at least one value.
func BuildBinaryHeap(mode BinaryHeapModeImpl) BinaryHeap {
	switch mode {
	case FastInsert:
		return &binaryHeapFastInsert{}
	case FastRemove:
		return &binaryHeapFastRemove{}
	case Equilibrated:
		return &binaryHeapBinaryTree{}
	}
	return nil
}

type binaryHeapFastInsert struct {
	a []int
}

// Insert appends without maintaining any order. Appending is amortized O(1),
// although growing the slice occasionally requires copying its values.
func (b *binaryHeapFastInsert) Insert(x int) {
	b.a = append(b.a, x)
}

// RemoveMin scans the unordered slice in O(n), moves the last value into the
// minimum's position, and shortens the slice without shifting other values.
func (b *binaryHeapFastInsert) RemoveMin() int {
	j := 0
	for i := range b.a {
		if b.a[i] < b.a[j] {
			j = i
		}
	}
	aux := b.a[j]
	b.a[j] = b.a[len(b.a)-1]
	b.a[len(b.a)-1] = aux

	b.a = b.a[:len(b.a)-1]
	return aux
}

type binaryHeapFastRemove struct {
	a []int
}

// Insert maintains the slice in descending order. In the worst case the new
// value crosses every existing value, making insertion O(n).
func (b *binaryHeapFastRemove) Insert(x int) {
	b.a = append(b.a, x)

	i := len(b.a) - 1
	for i > 0 && b.a[i] > b.a[i-1] {
		aux := b.a[i]
		b.a[i] = b.a[i-1]
		b.a[i-1] = aux
		i--
	}
}

// RemoveMin is O(1) because the descending order keeps the minimum at the end
// of the slice, where it can be removed by reslicing.
func (b *binaryHeapFastRemove) RemoveMin() int {
	aux := b.a[len(b.a)-1]
	b.a = b.a[:len(b.a)-1]
	return aux
}

type binaryHeapBinaryTree struct {
	h []int
}

// Insert appends the value at the next open position in the complete binary
// tree, then sifts it toward the root until the min-heap property is restored.
// With zero-based indexing, the parent of node i is (i-1)/2. The tree height,
// and therefore the worst-case insertion time, is O(log n).
func (b *binaryHeapBinaryTree) Insert(x int) {
	b.h = append(b.h, x)
	i := len(b.h) - 1

	for i > 0 && b.h[i] < b.h[(i-1)/2] {
		aux := b.h[i]
		b.h[i] = b.h[(i-1)/2]
		b.h[(i-1)/2] = aux

		i = (i - 1) / 2
	}
}

// RemoveMin removes the root, replaces it with the final leaf, and sifts that
// value down through its smaller child until the min-heap property is restored.
// The children of node i are 2*i+1 and 2*i+2, so removal is O(log n).
func (b *binaryHeapBinaryTree) RemoveMin() int {
	res := b.h[0]
	b.h[0] = b.h[len(b.h)-1]
	b.h = b.h[:len(b.h)-1]
	i := 0
	n := len(b.h)
	for 2*i+1 < n {
		j := 2*i + 1
		if (2*i+2 < n) && b.h[2*i+2] < b.h[j] {
			j = 2*i + 2
		}
		if b.h[j] >= b.h[i] {
			break
		}
		aux := b.h[j]
		b.h[j] = b.h[i]
		b.h[i] = aux
		i = j
	}

	return res
}
