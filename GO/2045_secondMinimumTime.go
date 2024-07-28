package main

import (
	"container/heap"
	"math"
)

// Pair represents a node and its corresponding travel time.
type Pair struct {
	node, time int
}

// PriorityQueue implements a priority queue for Pair elements.
type PriorityQueue []Pair

// Len returns the number of elements in the priority queue.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less reports whether the element with index i should sort before the element with index j.
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].time < pq[j].time }

// Swap exchanges the elements with indexes i and j.
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

// Push adds an element x to the priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Pair))
}

// Pop removes and returns the element with the highest priority (i.e., smallest time).
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// secondMinimum calculates the second minimum time to reach node n from node 1.
// Parameters:
// n: The number of nodes in the graph.
// edges: A list of edges where each edge is represented as [u, v].
// time: The time required to traverse an edge.
// change: The duration after which the traffic signal changes.
func secondMinimum(n int, edges [][]int, time int, change int) int {
	// Initialize the graph as an adjacency list.
	graph := make([][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Initialize the minTime array to track the first and second minimum times for each node.
	minTime := make([][2]int, n+1)
	for i := 1; i <= n; i++ {
		minTime[i][0] = math.MaxInt  // First minimum time
		minTime[i][1] = math.MaxInt  // Second minimum time
	}
	minTime[1][0] = 0 // Start node 1 with a time of 0

	// Initialize the priority queue with the start node.
	pq := &PriorityQueue{}
	heap.Push(pq, Pair{1, 0})

	// Process nodes from the priority queue.
	for pq.Len() > 0 {
		// Pop the node with the smallest time.
		current := heap.Pop(pq).(Pair)
		i, prevTime := current.node, current.time

		// Calculate the wait time based on the current time and signal change interval.
		numChangeSignal := prevTime / change
		waitTime := 0
		if numChangeSignal%2 != 0 {
			waitTime = change - prevTime%change
		}
		newTime := prevTime + waitTime + time

		// Update times for adjacent nodes.
		for _, j := range graph[i] {
			if newTime < minTime[j][0] {
				// If the new time is less than the first minimum time, update it.
				minTime[j][0] = newTime
				heap.Push(pq, Pair{j, newTime})
			} else if minTime[j][0] < newTime && newTime < minTime[j][1] {
				// If the new time is between the first and second minimum times, update the second minimum time.
				if j == n {
					return newTime // Return if we've reached the destination node n.
				}
				minTime[j][1] = newTime
				heap.Push(pq, Pair{j, newTime})
			}
		}
	}

	// Return -1 if the second minimum time is not found.
	return -1
}
