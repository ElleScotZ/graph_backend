package core

import (
	"sync"

	"github.com/ellescotz/graph_backend/pkg/utils"
)

// Node represents a node in the graph.
type Node struct {
	Name string `json:"name"`
}

// Edge represents an edge in the graph.
type Edge struct {
	Nodes  [2]Node `json:"nodes"` // cannot be identical
	Weight float64 `json:"weight"`
}

// Graph represents a directed graph with its nodes and edges.
type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"` // unique directed edges
}

// Path represents a subgraph in a directed graph along a path.
type Path struct {
	Subgraph Graph   `json:"subgraph"`
	Weight   float64 `json:"weight"` // sum weight of paths
}

// Equals checks equality between n and node2.
// It returns true if equal; otherwise false.
func (n *Node) Equals(node2 Node) bool {
	return n.Name == node2.Name
}

// Equals checks equality between e and edge2.
// It returns true if equal; otherwise false.
func (e *Edge) Equals(edge2 *Edge) bool {
	return e.Nodes == edge2.Nodes && e.Weight == edge2.Weight
}

// Copy returns a copy of g.
func (g *Graph) Copy() Graph {
	var copyGraph Graph

	copyGraph.Nodes = append(copyGraph.Nodes, g.Nodes...)
	copyGraph.Edges = append(copyGraph.Edges, g.Edges...)

	return copyGraph
}

// Copy returns a copy of P.
func (p *Path) Copy() Path {
	return Path{
		Subgraph: p.Subgraph.Copy(),
		Weight:   p.Weight}
}

// Weight adds up the weights of all edges along the path and returns it.
func (p *Path) GetWeight() float64 {
	p.Weight = 0

	for _, e := range p.Subgraph.Edges {
		p.Weight += e.Weight
	}

	return p.Weight
}

// findPathOnGraphWithoutEdgeRepetition generates subgraph
// containing all nodes starting from node1 ending with node2.
// It walks on g's nodes and edges without edge repetition.
// If a path is finished it puts it into paths map.
// It is a recursive function.
// excludeNodes: a map containing nodes to track path.
// path: pointer to the currently examined path.
// maxLength: max. number of nodes in path.
// maxWeight: max. sum weight of path.
func (g *Graph) findPathOnGraphWithoutEdgeRepetition(node1, node2 Node,
	excludeNodes *sync.Map, path *Path, paths *sync.Map, maxLength int, maxWeight float64) {
	// Appending node1 to path
	path.Subgraph.Nodes = append(path.Subgraph.Nodes, node1)

	// Excluding current node
	excludeNodes.Store(node1, true)

	var (
		wG        sync.WaitGroup
		stopLoop  = make(chan bool)
		pathFound bool
	)

	// Checking whether node1 is just one edge away from node2
	for i := 0; i < len(g.Edges) && !pathFound; i++ {
		wG.Add(1)

		go func(i int, stopLoop chan bool) {
			defer wG.Done()

			select {
			case _, ok := <-stopLoop:
				if !ok {
					return
				}
			default:
				if g.Edges[i].Nodes[0].Equals(node1) && g.Edges[i].Nodes[1].Equals(node2) {
					close(stopLoop)
					pathFound = true

					path.Subgraph.Nodes = append(path.Subgraph.Nodes, node2)
					path.Subgraph.Edges = append(path.Subgraph.Edges, g.Edges[i])

					if sumWeight := path.GetWeight(); sumWeight <= maxWeight || maxWeight == -1 {
						paths.Store(utils.SyncMapLength(*paths), path.Copy())

						// In order to check other possibilities, last element has to be removed.
						path.Subgraph.Nodes = path.Subgraph.Nodes[:len(path.Subgraph.Nodes)-1]
						path.Subgraph.Edges = path.Subgraph.Edges[:len(path.Subgraph.Edges)-1]
					}

					// Making sure new path is not equal to current one
					excludeNodes.Store(node2, true)
				}
			}
		}(i, stopLoop)
	}

	wG.Wait()

	if sumWeight := path.GetWeight(); (sumWeight <= maxWeight || maxWeight == -1) && len(path.Subgraph.Nodes) < maxLength {
		// Walking further from node1 on every nonexcluded edge in g.
		for i := range g.Edges {
			// wG.Add(1)

			// go func(i int) {
			// 	defer wG.Done()

			present, ok := excludeNodes.Load(g.Edges[i].Nodes[1])

			if g.Edges[i].Nodes[0].Equals(node1) && (!ok || present == false) {
				path.Subgraph.Edges = append(path.Subgraph.Edges, g.Edges[i])

				g.findPathOnGraphWithoutEdgeRepetition(g.Edges[i].Nodes[1], node2, excludeNodes, path, paths, maxLength, maxWeight)
			}
			// }(i)
		}
	}

	// wG.Wait()

	// If every adjacent node is excluded, or there is no edge from node1,
	// remove last node and edge from path.
	path.Subgraph.Nodes = path.Subgraph.Nodes[:len(path.Subgraph.Nodes)-1]

	if len(path.Subgraph.Edges) > 0 {
		path.Subgraph.Edges = path.Subgraph.Edges[:len(path.Subgraph.Edges)-1]
	}

	// Including node1 again
	excludeNodes.Store(node1, false)
}

// GeneratePathsWithoutEdgeRepetition finds all paths from node1 to node2.
// Each path contains an edge only once, no repetition is allowed.
// It returns a map of all paths.
func (g *Graph) GeneratePathsWithoutEdgeRepetition(node1, node2 Node) []Path {
	var (
		checkedNodes = sync.Map{} // to avoid repetition
		paths        = sync.Map{} // contains all paths of g from node1 to node2
		path         Path         // current path
		pathSlice    = make([]Path, 0, utils.SyncMapLength(paths))
	)

	g.findPathOnGraphWithoutEdgeRepetition(node1, node2, &checkedNodes, &path, &paths, len(g.Nodes), -1)

	paths.Range(func(key, value interface{}) bool {
		path := value.(Path)

		pathSlice = append(pathSlice, path)

		return true
	})

	return pathSlice
}

// GeneratePthsWithMaxSteps finds paths from node1 to node2 in no more than numberOfEdges steps.
// Edge repetition is not allowed.
// numberOfEdges: maximum amount of edges required for paths.
// exactSteps: true, if paths have to contain exactly numberOfEdges steps; otherwise false.
func (g *Graph) GeneratePathsWithMaxSteps(node1, node2 Node, numberOfEdges int, exactSteps bool) []Path {
	var (
		checkedNodes = sync.Map{} // to avoid repetition
		paths        = sync.Map{} // contains all paths of g from node1 to node2
		path         Path         // current path
		pathSlice    = make([]Path, 0, utils.SyncMapLength(paths))
	)

	g.findPathOnGraphWithoutEdgeRepetition(node1, node2, &checkedNodes, &path, &paths, numberOfEdges, -1)

	paths.Range(func(key, value interface{}) bool {
		path := value.(Path)

		if (exactSteps && len(path.Subgraph.Nodes) == numberOfEdges+1) || !exactSteps {
			pathSlice = append(pathSlice, path)
		}

		return true
	})

	return pathSlice
}

// GeneratePathsWithMaxWeight finds paths from node1 to node2 with no greater weight than sumWeight.
// Edge repetition is not allowed.
// sumWeight: maximum sum of weight along the paths.
// exactWeight: true, if paths can have exactly sumWeight; otherwise false.
func (g *Graph) GeneratePathsWithMaxWeight(node1, node2 Node, sumWeight float64, exactWeight bool) []Path {
	var (
		checkedNodes = sync.Map{} // to avoid repetition
		paths        = sync.Map{} // contains all paths of g from node1 to node2
		path         Path         // current path
		pathSlice    = make([]Path, 0, utils.SyncMapLength(paths))
	)

	g.findPathOnGraphWithoutEdgeRepetition(node1, node2, &checkedNodes, &path, &paths, len(g.Nodes), sumWeight)

	paths.Range(func(key, value interface{}) bool {
		path := value.(Path)

		if (exactWeight && path.Weight == sumWeight) || !exactWeight {
			pathSlice = append(pathSlice, path)
		}

		return true
	})

	return pathSlice
}

// GenerateLowestHighestWeightPath finds the path from node1 to node2 that is
// either the lowest or highest weighted path.
// lowest: true, if path is the lowest weighted; otherwise false.
// Naive solution to Travelling Salesperson Problem and Hamiltonian Cycle Problem.
func (g *Graph) GenerateLowestHighestWeightPath(node1, node2 Node, lowest bool) []Path {
	paths := g.GeneratePathsWithoutEdgeRepetition(node1, node2)

	if len(paths) == 0 {
		return []Path{}
	}

	var (
		weight        = paths[0].Weight
		relevantPaths []Path
	)

	for i := range paths {
		if lowest && paths[i].Weight <= weight {
			if paths[i].Weight < weight {
				relevantPaths = nil
			}

			weight = paths[i].Weight
			relevantPaths = append(relevantPaths, paths[i])
		} else if !lowest && paths[i].Weight >= weight {
			if paths[i].Weight > weight {
				relevantPaths = nil
			}

			weight = paths[i].Weight
			relevantPaths = append(relevantPaths, paths[i])
		}
	}

	return relevantPaths
}

// GenerateShortestLongestPath find the shortest/longest path from node1 to node2.
// shortest: true, if path is the shortest; otherwise false.
func (g *Graph) GenerateShortestLongestPath(node1, node2 Node, shortest bool) []Path {
	paths := g.GeneratePathsWithoutEdgeRepetition(node1, node2)

	if len(paths) == 0 {
		return []Path{}
	}

	var (
		length        = len(paths[0].Subgraph.Nodes)
		relevantPaths []Path
	)

	for i := range paths {
		if shortest && len(paths[i].Subgraph.Nodes) <= length {
			if len(paths[i].Subgraph.Nodes) < length {
				relevantPaths = nil
			}

			length = len(paths[i].Subgraph.Nodes)
			relevantPaths = append(relevantPaths, paths[i])
		} else if !shortest && len(paths[i].Subgraph.Nodes) >= length {
			if len(paths[i].Subgraph.Nodes) > length {
				relevantPaths = nil
			}

			length = len(paths[i].Subgraph.Nodes)
			relevantPaths = append(relevantPaths, paths[i])
		}
	}

	return relevantPaths
}
