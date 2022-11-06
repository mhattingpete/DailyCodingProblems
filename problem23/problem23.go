package main

import (
    "fmt"
    "example.com/graph"
)

/*
You are given an M by N matrix consisting of booleans that represents a board. 
Each True boolean represents a wall. Each False boolean represents a tile you can walk on.

Given this matrix, a start coordinate, and an end coordinate, return the minimum number of steps required to reach the end coordinate from the start. 
If there is no possible path, then return null. You can move up, left, down, and right. You cannot move through walls. You cannot wrap around the edges of the board.

For example, given the following board:

[[f, f, f, f],
[t, t, f, t],
[f, f, f, f],
[f, f, f, f]]
and start = (3, 0) (bottom left) and end = (0, 0) (top left), the minimum number of steps required to reach the end is 7, 
since we would need to go through (1, 2) because there is a wall everywhere else on the second row.


Nodes:
0	1	2	3
4	5	6	7
8	9	10	11
12	13	14	15
*/


func ConstructGraph(matrix [][]bool) *graph.Graph {
	g := &graph.Graph{}
	for i := 0; i < len(matrix)*len(matrix[0]); i++ {
		g.AddNode(&graph.Node{Value:i})
	}
	n := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == false {
				if j+1 < len(matrix[0]) {
					if matrix[i][j+1] == false {
						g.AddEdge(g.Nodes[n], g.Nodes[n+1])
					}
				} 
				if j-1 > 0 {
					if matrix[i][j-1] == false {
						g.AddEdge(g.Nodes[n], g.Nodes[n-1])
					}
				}
				if i+1 < len(matrix) {
					if matrix[i+1][j] == false {
						g.AddEdge(g.Nodes[n], g.Nodes[n+len(matrix[0])])
					}
				} 
				if i-1 > 0 {
					if matrix[i-1][j] == false {
						g.AddEdge(g.Nodes[n], g.Nodes[n-len(matrix[0])])
					}
				}
			}
			n++
		}
	}
	return g
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func Backtrace(parent map[int]*graph.Node, start int, end int) []int {
	var path []int
	var parentNode *graph.Node
	path = append(path, end)
	for end != start {
		parentNode = parent[end]
		path = append(path, parentNode.Value)
		end = parentNode.Value
	}
	reverse(path)
	return path
}

func BreadthFirstSearch(g *graph.Graph, start int, end int) []int { // Special case of Dijkstras Algorithm as the g is unweighted
	parent := make(map[int]*graph.Node)
	q := graph.NodeQueue{}
	n := g.Nodes[start]
	q.Enqueue(*n)
	visited := make(map[*graph.Node]bool)
	for q.IsEmpty() == false {
		node := q.Dequeue()
		if node.Value == end {
			return Backtrace(parent, start, end)
		}
		visited[node] = true
		near := g.Edges[*node]
		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				parent[j.Value] = node
				q.Enqueue(*j)
				visited[j] = true
			}
		}
	}
	return nil
}

func main() {
	matrix := [][]bool{[]bool{false, false, false, false}, []bool{true, true, false, true}, []bool{false, false, false, false}, []bool{false, false, false, false}}
	g := ConstructGraph(matrix)
	fmt.Println(g.To_String())
	startX := 3
	startY := 3
	endX := 0
	endY := 0
	start := startX + (len(matrix[0])*startY)
	end := endX + (len(matrix[0])*endY)
	path := BreadthFirstSearch(g, start, end)
	fmt.Println(path)
}