package tree

type Tree struct {
	Data       string
	Childs     []Tree
	discovered bool
}

// BFS is the implementation of Breadth-first search
func (t *Tree) BFS() (res []string) {
	q := []Tree{}

	t.discovered = true
	res = append(res, t.Data)
	q = append(q, *t)
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, tree := range v.Childs {
			if !tree.discovered {
				res = append(res, tree.Data)
				tree.discovered = true
				q = append(q, tree)
			}
		}
	}
	return res
}

// BFSDepthLimit is the implementation of Breadth-first search but with a limit
// in depth level, n represents de max depth level
func (t *Tree) BFSDepthLimit(n int) (res []string) {
	depthCount := 0
	q := []Tree{}

	t.discovered = true
	res = append(res, t.Data)
	q = append(q, *t)
	for len(q) > 0 && depthCount < n {
		v := q[0]
		q = q[1:]
		for _, tree := range v.Childs {
			if !tree.discovered {
				res = append(res, tree.Data)
				tree.discovered = true
				q = append(q, tree)
			}
		}
		depthCount++
	}
	return res
}
