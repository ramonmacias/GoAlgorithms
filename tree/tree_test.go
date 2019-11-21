package tree_test

import (
	"reflect"
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
	want := []string{"root", "child1", "child2", "child3", "child10", "child11"}
	levelOne := []tree.Tree{tree.Tree{Data: "child1", Childs: buildNodes(2, "child1")}, tree.Tree{Data: "child2"}, tree.Tree{Data: "child3"}}
	tree := tree.Tree{
		Data:   "root",
		Childs: levelOne,
	}
	got := tree.BFS()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("We want %+v but got %+v", want, got)
	}
}

func TestBFSDepthLimit(t *testing.T) {
	want := []string{"root", "child1", "child2", "child3"}
	levelOne := []tree.Tree{tree.Tree{Data: "child1", Childs: buildNodes(2, "child1")}, tree.Tree{Data: "child2"}, tree.Tree{Data: "child3"}}
	tree := tree.Tree{
		Data:   "root",
		Childs: levelOne,
	}
	got := tree.BFSDepthLimit(1)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("We want %+v but got %+v", want, got)
	}
}
