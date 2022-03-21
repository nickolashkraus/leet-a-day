package pkg

import (
	"reflect"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	adjList := [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}
	nodes := graphFromAdjacencyList(adjList)
	// Pointer to the first node in the graph representing the adjacency list
	// 'adjList'.
	node := nodes[0]
	// Pointer to the node in the deep copy of the graph.
	nodeCopy := cloneGraph(node)
	newAdjList := &[][]int{}
	nodeToAdjacencyList(nodeCopy, &[]*Node{}, newAdjList)
	// Check if the adjacency list generated from the copied node in the graph
	// representing the original adjacency list is the same as the original.
	//
	// Result:
	//   * 'adjList' and 'newAdjList' should be identical.
	if !reflect.DeepEqual(adjList, *newAdjList) {
		t.Fail()
	}
	newNodes := graphFromAdjacencyList(*newAdjList)
	// Check if new nodes were copied.
	//
	// Result:
	//   * Each node in 'newNodes' should be a copy of the node in 'nodes'. That
	//     is, they should be identical, but not the same node.
	for i := range nodes {
		if nodes[i] == newNodes[i] {
			t.Fail()
		}
		if !reflect.DeepEqual(nodes[i], newNodes[i]) {
			t.Fail()
		}
	}
}

func TestIsVisited(t *testing.T) {
	// Check if pointer to 'node' is in map 'visited'.
	//
	// Result:
	//   * 'visited' is allocated with pointer to 'node', therefore isVisited
	//     should return true.
	node := &Node{
		Val:       0,
		Neighbors: []*Node{},
	}
	visited := []*Node{node}
	if !isVisited(node, visited) {
		t.Fail()
	}
	// Check if pointer to 'nodeCopy' is not in map 'visited'.
	//
	// Result:
	//   * 'nodeCopy' is not added to 'visited', therefore isVisited should
	//     return false.
	nodeCopy := &Node{
		Val:       0,
		Neighbors: []*Node{},
	}
	if isVisited(nodeCopy, visited) {
		t.Fail()
	}
}

func TestGraphFromAdjacencyList(t *testing.T) {
	adjList := [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}
	nodes := graphFromAdjacencyList(adjList)
	newAdjList := [][]int{}
	// For a slice of structs of type Node, extract the value (Node.Val) of its
	// neighbors and create a new adjacency list.
	for _, node := range nodes {
		neighbors := []int{}
		for _, neighbor := range node.Neighbors {
			neighbors = append(neighbors, neighbor.Val)
		}
		newAdjList = append(newAdjList, neighbors)
	}
	// Check if the adjacency list generated from the nodes of the original
	// adjacency list is the same as the original.
	//
	// Result:
	//   * 'adjList' and 'newAdjList' should be identical.
	if !reflect.DeepEqual(adjList, newAdjList) {
		t.Fail()
	}
}

func TestNodeToAdjacencyList(t *testing.T) {
	adjList := [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}
	nodes := graphFromAdjacencyList(adjList)
	newAdjList := &[][]int{}
	nodeToAdjacencyList(nodes[0], &[]*Node{}, newAdjList)
	// Check if the adjacency list generated from a node in a graph representing
	// the original adjacency list is the same as the original.
	//
	// Result:
	//   * 'adjList' and 'newAdjList' should be identical.
	if !reflect.DeepEqual(adjList, *newAdjList) {
		t.Fail()
	}
}
