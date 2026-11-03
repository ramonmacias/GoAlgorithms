package structures

// BinaryHeap describes the two essential operations of a min-priority queue:
// adding a value and removing the smallest value currently stored.
//
// The binary heap was introduced by J. W. J. Williams in his 1964 paper
// "Algorithm 232: Heapsort." Williams represented a complete binary tree in an
// array, avoiding pointers, and described operations for inserting a value and
// extracting the minimum. A conventional binary heap balances both operations
// at O(log n). The implementations below instead demonstrate the two simpler
// extremes: make either insertion or removal fast at the other's expense.
type BinaryHeap interface {
	Insert(x int)
	RemoveMin() int
}

// BuildBinaryHeap returns a slice-backed min-priority queue. If fastInsert is
// true, insertion is amortized O(1) and removal is O(n). Otherwise, insertion
// is O(n) and removal is O(1).
//
// RemoveMin must only be called when the queue contains at least one value.
func BuildBinaryHeap(fastInsert bool) BinaryHeap {
	if fastInsert {
		return &binaryHeapFastInsert{}
	}
	return &binaryHeapFastRemove{}
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
