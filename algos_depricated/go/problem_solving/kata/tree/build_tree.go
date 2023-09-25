package tree

import (
	"fmt"
	"strings"
)

// Node is a node in a tree
type Node struct {
	ID       int
	Name     string
	parentID int
	Children []*Node
}

var flatData = []Node{
	{ID: 1, Name: "Europe", parentID: 0},
	{ID: 2, Name: "Asia", parentID: 0},
	{ID: 3, Name: "Germany", parentID: 1},
	{ID: 4, Name: "France", parentID: 1},
	{ID: 5, Name: "Italy", parentID: 1},
	{ID: 6, Name: "Spain", parentID: 1},
	{ID: 7, Name: "England", parentID: 1},
	{ID: 8, Name: "India", parentID: 2},
	{ID: 9, Name: "China", parentID: 2},
	{ID: 10, Name: "Japan", parentID: 2},
	{ID: 11, Name: "Berlin", parentID: 3},
	{ID: 12, Name: "Munich", parentID: 3},
	{ID: 13, Name: "Hamburg", parentID: 3},
	{ID: 14, Name: "Paris", parentID: 4},
	{ID: 15, Name: "Marseille", parentID: 4},
	{ID: 16, Name: "Lyon", parentID: 4},
	{ID: 17, Name: "Rome", parentID: 5},
	{ID: 18, Name: "Milan", parentID: 5},
	{ID: 19, Name: "Naples", parentID: 5},
	{ID: 20, Name: "Madrid", parentID: 6},
	{ID: 21, Name: "Barcelona", parentID: 6},
	{ID: 22, Name: "Valencia", parentID: 6},
	{ID: 23, Name: "London", parentID: 7},
	{ID: 24, Name: "Delhi", parentID: 8},
	{ID: 25, Name: "Beijing", parentID: 9},
	{ID: 26, Name: "Tokyo", parentID: 10},
	{ID: 27, Name: "Poland", parentID: 1},
	{ID: 28, Name: "Warsaw", parentID: 27},
	{ID: 29, Name: "Krakow", parentID: 27},
	{ID: 30, Name: "Wroclaw", parentID: 27},
	{ID: 31, Name: "Gdansk", parentID: 27},
	{ID: 32, Name: "Lodz", parentID: 27},
	{ID: 33, Name: "Bialystok", parentID: 27},
	{ID: 34, Name: "Mokotow", parentID: 28},
	{ID: 35, Name: "Praga", parentID: 28},
	{ID: 36, Name: "Ursus", parentID: 28},
	{ID: 37, Name: "Zoliborz", parentID: 28},
}

func buildTree(data []Node, parentID int) []*Node {
	var nodes []Node
	for _, x := range data {
		if x.parentID == parentID {
			nodes = append(nodes, x)
		}
	}
	if len(nodes) == 0 {
		return nil
	}
	var result []*Node
	for _, x := range nodes {
		n := &Node{ID: x.ID,
			Name:     x.Name,
			parentID: x.parentID,
			Children: buildTree(data, x.ID)}
		result = append(result, n)
	}
	return result
}

func printTree(tree []*Node, level int) {
	for _, x := range tree {
		fmt.Println(strings.Repeat("  ", level), x.Name)
		printTree(x.Children, level+1)
	}
}
