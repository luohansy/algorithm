package main

import (
   "github.com/luohansy/algorithm/libgo"
   "fmt"
)

func main() {
	vertexes := []libgo.GVertex{"v1", "v2","v3","v4", "v5", "v6", "v7"}
	edges := []libgo.GEdge {
		libgo.GEdge{"v1", "v2", 20},
		libgo.GEdge{"v2", "v3", 15},
		libgo.GEdge{"v3", "v4", 3},
		libgo.GEdge{"v4", "v5", 17},
		libgo.GEdge{"v5", "v6", 28},
		libgo.GEdge{"v6", "v1", 23},
		libgo.GEdge{"v7", "v1", 1},
		libgo.GEdge{"v7", "v2", 4},
		libgo.GEdge{"v7", "v3", 9},
		libgo.GEdge{"v7", "v4", 16},
		libgo.GEdge{"v7", "v5", 25},
		libgo.GEdge{"v7", "v6", 36},
	}
	ptg := libgo.NewGraph(vertexes, edges)
	fmt.Printf("%v\n", ptg)
}
