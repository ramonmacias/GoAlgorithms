package tree

import "log"

type Tree struct {
	Data       string
	Childs     []Tree
	discovered bool
}

// BFS is the implementation of Breadth-first search
func (t *Tree) BFS() {
	q := []Tree{}

	t.discovered = true
	log.Println(t.Data)
	q = append(q, *t)
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, tree := range v.Childs {
			if !tree.discovered {
				log.Println(tree.Data)
				tree.discovered = true
				q = append(q, tree)
			}
		}
	}
}
