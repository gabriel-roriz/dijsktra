package dijkstra

import "math"

func initialize(g *Graph, origin *Vertice) {

	for _, Vertice := range g.Vertices {

		if Vertice == origin {
			(*Vertice).Distance = 0
		} else {
			(*Vertice).Distance = math.Inf(99999)
		}

		(*Vertice).Predecessor = nil
	}

}

func relax(from *Vertice, Edge Edge) {

	to := Edge.To

	if (*to).Distance > ((*from).Distance + Edge.Weight) {
		(*to).Distance = (*from).Distance + Edge.Weight
		(*to).Predecessor = from
	}
}

func findNext(Graph *Graph, processeds []*Vertice) *Vertice {

	var lowest *Vertice

Graph_LOOP:
	for _, Vertice := range Graph.Vertices {

		//if is already processed, continue
		for _, processed := range processeds {
			if processed == Vertice {
				continue Graph_LOOP
			}
		}

		if lowest != nil {
			//find the first one who isn't in processed list
			if (*Vertice).Distance < (*lowest).Distance {
				lowest = Vertice
			}

		} else {
			//find the first one who isn't in processed list
			lowest = Vertice
		}
	}

	return lowest
}

func Run(Graph *Graph, first *Vertice) {
	initialize(Graph, first)

	processed := []*Vertice{}

	node := first

	for node != nil {

		for _, Edge := range node.Edges {
			relax(node, Edge)
		}
		processed = append(processed, node)
		node = findNext(Graph, processed)
	}

}
