package dijkstra

type Vertice struct {
	Label string
	Edges []Edge

	Distance    float64
	Predecessor *Vertice
}

type Edge struct {
	Weight float64
	To     *Vertice
}

type Graph struct {
	Vertices []*Vertice
}
