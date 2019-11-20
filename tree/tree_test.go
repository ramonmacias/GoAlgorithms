package tree_test

import (
	"strconv"
	"testing"

	"github.com/ramonmacias/GoAlgorithms/tree"
)

func buildNodes(size int, suffix string) []tree.Tree {
	nodes := []tree.Tree{}
	for i := 0; i < size; i++ {
		nodes = append(nodes, tree.Tree{Data: suffix + strconv.Itoa(i)})
	}
	return nodes
}

func TestBFS(t *testing.T) {
	levelOne := []tree.Tree{tree.Tree{Data: "child1", Childs: buildNodes(2, "child1")}, tree.Tree{Data: "child2"}, tree.Tree{Data: "child3"}}
	tree := tree.Tree{
		Data:   "root",
		Childs: levelOne,
	}
	tree.BFS()
}
