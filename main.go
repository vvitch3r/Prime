package main

import (
	"container/heap"
	"fmt"
)

// Edge represents an edge in the graph
type Edge struct {
	to, weight int
}

// Graph represents a weighted undirected graph using an adjacency list
type Graph struct {
	vertices int
	adjList  [][]Edge
}

// NewGraph initializes a new graph with a given number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make([][]Edge, vertices),
	}
}

// AddEdge adds an undirected edge to the graph
func (g *Graph) AddEdge(from, to, weight int) {
	g.adjList[from] = append(g.adjList[from], Edge{to, weight})
	g.adjList[to] = append(g.adjList[to], Edge{from, weight})
}

// Item represents an item in the priority queue
type Item struct {
	vertex, weight int
}

// PriorityQueue implements a min-heap for Items
type PriorityQueue []*Item

// Len returns the length of the priority queue
func (pq PriorityQueue) Len() int { return len(pq) }

// Less compares two items in the priority queue
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

// Swap swaps two items in the priority queue
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push adds an item to the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

// Pop removes and returns the smallest item from the priority queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// PrimMST finds the Minimum Spanning Tree using Prim's algorithm
func (g *Graph) PrimMST() int {
	// Total weight of the MST
	totalWeight := 0

	// Visited array to keep track of visited vertices
	visited := make([]bool, g.vertices)

	// Priority queue to select the edge with the minimum weight
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Start from vertex 0 (can start from any vertex)
	heap.Push(pq, &Item{vertex: 0, weight: 0})

	for pq.Len() > 0 {
		// Get the edge with the smallest weight
		item := heap.Pop(pq).(*Item)
		vertex := item.vertex
		weight := item.weight

		// If the vertex is already visited, skip it
		if visited[vertex] {
			continue
		}

		// Mark the vertex as visited
		visited[vertex] = true

		// Add the weight to the total weight of the MST
		totalWeight += weight

		// Add all adjacent edges to the priority queue
		for _, edge := range g.adjList[vertex] {
			if !visited[edge.to] {
				heap.Push(pq, &Item{vertex: edge.to, weight: edge.weight})
			}
		}
	}

	return totalWeight
}

func main() {
	// Create a new graph with 5 vertices
	graph := NewGraph(5)

	// Add edges to the graph
	graph.AddEdge(0, 1, 2)
	graph.AddEdge(0, 3, 6)
	graph.AddEdge(1, 2, 3)
	graph.AddEdge(1, 3, 8)
	graph.AddEdge(1, 4, 5)
	graph.AddEdge(2, 4, 7)
	graph.AddEdge(3, 4, 9)

	// Find the MST using Prim's algorithm
	totalWeight := graph.PrimMST()

	// Print the total weight of the MST
	fmt.Printf("Total weight of the Minimum Spanning Tree: %d\n", totalWeight)
}
