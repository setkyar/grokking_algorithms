package main

import (
	"fmt"
	"math"
)

// Dijkstra's algorithm Algorithm
// use when it's finding shortest path on weighted graphs
// Won't work on graph cycle
func main() {
	// create a hash table to store graph
	// [start] -> [a:6, b:2]
	graph := make(map[string]map[string]float64)
	graph["start"] = make(map[string]float64)
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	// [a] -> [fin:1]
	graph["a"] = make(map[string]float64)
	graph["a"]["fin"] = 1

	// [b] -> [a:3, fin:5]
	graph["b"] = make(map[string]float64)
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5

	// [fin] -> []
	graph["fin"] = make(map[string]float64)

	// create a hash table to store the costs for each node
	// [fin] -> infinity
	infinity := math.Inf(1) // If you don't know -> https://www.educative.io/answers/what-is-the-inf-function-in-golang

	// create cost table
	costs := make(map[string]float64)

	// [a] -> 6
	costs["a"] = 6

	// [b] -> 2
	costs["b"] = 2
	costs["fin"] = infinity

	// parents table
	parents := make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""

	// processed nodes
	processed := make([]string, 0)

	node := find_lowest_cost_node(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for n := range neighbors {
			newCost := cost + neighbors[n]
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}

		// mark the node as processed
		processed = append(processed, node)
		// find the next node to process, and loop
		node = find_lowest_cost_node(costs, processed)
	}

	// print the result
	fmt.Println(costs)
}

func find_lowest_cost_node(costs map[string]float64, processed []string) string {
	lowest_cost := math.Inf(1)
	lowest_cost_node := ""
	for node := range costs {
		cost := costs[node]
		if cost < lowest_cost && !contains(processed, node) {
			lowest_cost = cost
			lowest_cost_node = node
		}
	}

	return lowest_cost_node
}

// check string contains in slice
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
