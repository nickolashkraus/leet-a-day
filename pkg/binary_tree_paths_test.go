package pkg

import (
	"reflect"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	tree := []string{"1", "2", "3", "null", "5"}
	// Pointer to the root node of the binary tree.
	rootNode := binaryTreeFromSlice(tree, 0)
	ret := binaryTreePaths(rootNode)
	if !reflect.DeepEqual([]string{"1->2->5", "1->3"}, ret) {
		t.Fail()
	}
	tree = []string{"1"}
	// Pointer to the root node of the binary tree.
	rootNode = binaryTreeFromSlice(tree, 0)
	ret = binaryTreePaths(rootNode)
	if !reflect.DeepEqual([]string{"1"}, ret) {
		t.Fail()
	}
}

func TestBinaryTreeFromSlice(t *testing.T) {
	tree := []string{"1", "2", "3", "null", "5"}
	// Pointer to the root node of the binary tree.
	rootNode := binaryTreeFromSlice(tree, 0)
	// Check binary tree:
	//
	// [1,2,3,null,5]
	//
	//     1
	//   /   \
	//  2     3
	//   \
	//     5
	if rootNode.Val != 1 {
		t.Fail()
	}
	if rootNode.Left.Val != 2 {
		t.Fail()
	}
	if rootNode.Right.Val != 3 {
		t.Fail()
	}
	if rootNode.Left.Right.Val != 5 {
		t.Fail()
	}
}
