package libgo

import (
//   "container/heap"
)

// graphic edge definition
type GEdge struct {
	EndpointX GVertex
	EndpointY GVertex
	Weight int
}

type GVertex string

// adjacent list graphic vertex definition
type adjVertex struct {
	vertex GVertex
	adjacencies []GEdge
}

// undirected graph definition
type Graph struct {
	adjTable []adjVertex
}

func NewGraph(vertexes []GVertex, edges []GEdge) (graph *Graph) {
	graph = &Graph{make([]adjVertex, 0, len(vertexes))}
	for _, v := range vertexes {
		adjEntry := adjVertex{v, make([]GEdge, 0, len(vertexes))}
		for _, e := range edges {
			if e.EndpointX == v || e.EndpointY == v {
				adjEntry.adjacencies = append(adjEntry.adjacencies, e)
			}
		}
		graph.adjTable = append(graph.adjTable, adjEntry)
	}

	return
}
