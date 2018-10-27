package main

import (
	"math"
)

type vertice struct {
	label string
	edges []edge

	distance    float64
	predecessor *vertice
}

type edge struct {
	weight float64
	to     *vertice
}

type graph struct {
	vertices []*vertice
}

func initialize(g *graph, origin *vertice) {

	for _, vertice := range g.vertices {

		if vertice == origin {
			(*vertice).distance = 0
		} else {
			(*vertice).distance = math.Inf(99999)
		}

		(*vertice).predecessor = nil
	}

}

func relax(from *vertice, edge edge) {

	to := edge.to

	if (*to).distance > ((*from).distance + edge.weight) {
		(*to).distance = (*from).distance + edge.weight
		(*to).predecessor = from
	}
}

func findNext(graph *graph, processeds []*vertice) *vertice {

	var lowest *vertice

GRAPH_LOOP:
	for _, vertice := range graph.vertices {

		//if is already processed, continue
		for _, processed := range processeds {
			if processed == vertice {
				continue GRAPH_LOOP
			}
		}

		if lowest != nil {
			//find the first one who isn't in processed list
			if (*vertice).distance < (*lowest).distance {
				lowest = vertice
			}

		} else {
			//find the first one who isn't in processed list
			lowest = vertice
		}
	}

	return lowest
}

func main() {

	a := vertice{label: "A"}
	b := vertice{label: "B"}
	c := vertice{label: "C"}
	d := vertice{label: "D"}
	e := vertice{label: "E"}
	f := vertice{label: "F"}

	a.edges = []edge{
		{to: &b, weight: 5},
		{to: &c, weight: 2}}

	b.edges = []edge{
		{to: &d, weight: 4},
		{to: &e, weight: 2}}

	c.edges = []edge{
		{to: &b, weight: 8},
		{to: &e, weight: 7}}

	d.edges = []edge{
		{to: &e, weight: 6},
		{to: &f, weight: 3}}

	e.edges = []edge{{to: &f, weight: 1}}

	f.edges = []edge{}

	graph := graph{vertices: []*vertice{&a, &b, &c, &d, &e, &f}}

	initialize(&graph, &a)

	processed := []*vertice{}

	node := &a

	for node != nil {

		for _, edge := range node.edges {
			relax(node, edge)
		}
		processed = append(processed, node)
		node = findNext(&graph, processed)
	}

}
