package main

// 133. Clone Graph
//
// https://leetcode.com/problems/clone-graph
//
// NOTES
//   * A connected graph is a graph for which there is a path from any point to
//     any other point in the graph.
//   * An undirected graph is a graph for which the relations between pairs of
//     vertices are symmetric, so that each edge has no directional character.
//   * An adjacency list is a collection of unordered lists used to represent a
//     finite graph. Each list describes the set of neighbors of a node in the
//     graph.

/**
 * Definition for a Node.
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	} else if len(node.Neighbors) == 0 {
		return &Node{
			Val:       node.Val,
			Neighbors: []*Node{},
		}
	} else {
		// Maps are reference data types, therefore they are passed by reference to
		// functions.
		visited := make(map[*Node]*Node)
		return copyNode(node, visited)
	}
}

func copyNode(node *Node, visited map[*Node]*Node) *Node {

	newNode, _ := visited[node]
	if newNode != nil {
		return newNode
	}

	nodeCopy := &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}
	visited[node] = nodeCopy

	for _, neighbor := range node.Neighbors {
		newNeighbor, _ := visited[neighbor]
		if newNeighbor != nil {
			nodeCopy.Neighbors = append(nodeCopy.Neighbors, newNeighbor)
		} else {
			nodeCopy.Neighbors = append(nodeCopy.Neighbors, copyNode(neighbor, visited))
		}
	}

	return nodeCopy
}

func isVisited(node *Node, visited []*Node) bool {
	for i := 0; i < len(visited); i++ {
		if node == visited[i] {
			return true
		}
	}
	return false
}

// Given an adjacency list, return a slice of all its nodes.
func graphFromAdjacencyList(adjList [][]int) []*Node {

	nodes := []*Node{}

	for i := range adjList {
		nodes = append(nodes, &Node{
			Val:       i + 1,
			Neighbors: []*Node{},
		})
	}

	for i, neighbors := range adjList {
		for _, neighbor := range neighbors {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[neighbor-1])
		}
	}

	return nodes
}

// Given a reference to a node in a connected undirected graph, generate an
// adjacency list representing the graph.
func nodeToAdjacencyList(node *Node, visited *[]*Node, adjList *[][]int) {

	*visited = append(*visited, node)

	neighbors := []int{}
	for _, neighbor := range node.Neighbors {
		neighbors = append(neighbors, neighbor.Val)
	}

	*adjList = append(*adjList, neighbors)

	for _, neighbor := range node.Neighbors {
		if !isVisited(neighbor, *visited) {
			nodeToAdjacencyList(neighbor, visited, adjList)
		}
	}

	return
}
