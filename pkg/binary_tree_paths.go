package pkg

import (
	"strconv"
)

// 257. Binary Tree Paths
//
// https://leetcode.com/problems/binary-tree-paths
//
// NOTES
//   * Use depth-first search to construct paths.
//
//     See the Wikipedia article:
//     https://en.wikipedia.org/wiki/Depth-first_search

/**
 * Definition for a TreeNode.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	paths := [][]*TreeNode{}
	pathStrings := []string{}

	if root.Left == nil && root.Right == nil {
		paths = append(paths, []*TreeNode{root})
	}

	if root.Left != nil {
		getPaths(root.Left, &paths, []*TreeNode{root})
	}

	if root.Right != nil {
		getPaths(root.Right, &paths, []*TreeNode{root})
	}

	for _, path := range paths {
		str := ""
		for i, node := range path {
			str += strconv.Itoa(node.Val)
			if i != len(path)-1 {
				str += "->"
			}
		}
		pathStrings = append(pathStrings, str)
	}

	return pathStrings
}

func getPaths(node *TreeNode, paths *[][]*TreeNode, path []*TreeNode) {
	path = append(path, node)

	if node.Left == nil && node.Right == nil {
		tmp := make([]*TreeNode, len(path))
		copy(tmp, path)
		*paths = append(*paths, tmp)
	}

	if node.Left != nil {
		getPaths(node.Left, paths, path)
	}

	if node.Right != nil {
		getPaths(node.Right, paths, path)
	}

	return
}

// Given a slice representing a binary tree, return a pointer to its root.
//
// For a binary tree, the following applies:
//   * parent = i
//   * left   = i*2 + 1
//   * right  = i*2 + 2
//
// Where i is the current index of the slice.
//
// Example:
//
//   [1,2,3,null,5]
//
//        1
//      /   \
//     2     3
//      \
//        5
//
func binaryTreeFromSlice(slice []string, index int) *TreeNode {
	// If null, return empty pointer (nil)
	if slice[index] == "null" {
		return nil
	}

	val, _ := strconv.Atoi(slice[index])
	node := TreeNode{
		Val: val,
	}

	if index*2+1 > len(slice)-1 {
		node.Left = nil
	} else {
		node.Left = binaryTreeFromSlice(slice, index*2+1)
	}

	if index*2+2 > len(slice)-1 {
		node.Right = nil
	} else {
		node.Right = binaryTreeFromSlice(slice, index*2+2)
	}

	return &node
}
