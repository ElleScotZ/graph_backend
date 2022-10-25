package core

import (
	"testing"
)

func TestEquals(t *testing.T) {
	t.Parallel()

	nodeA := Node{Name: "A"}
	nodeB := Node{Name: "B"}
	nodeC := Node{Name: "A"}
	nodeD := Node{Name: "D"}

	edgeAB := Edge{
		Nodes:  [2]Node{nodeA, nodeB},
		Weight: 2,
	}

	edgeBA := Edge{
		Nodes:  [2]Node{nodeB, nodeA},
		Weight: 2,
	}

	edgeBC := Edge{
		Nodes:  [2]Node{nodeB, nodeC},
		Weight: 2,
	}

	edgeCD := Edge{
		Nodes:  [2]Node{nodeC, nodeD},
		Weight: 2,
	}

	edgeAB2 := Edge{
		Nodes:  [2]Node{nodeA, nodeB},
		Weight: 4,
	}

	if edgeAB.Equals(&edgeBA) {
		t.Errorf("Equal did not work. Got true instead of false")
	}

	if !edgeBA.Equals(&edgeBC) {
		t.Errorf("Equal did not work. Got false instead of true")
	}

	if edgeAB.Equals(&edgeCD) {
		t.Errorf("Equal did not work. Got true instead of false")
	}

	if edgeAB.Equals(&edgeAB2) {
		t.Errorf("Equal did not work. Got true instead of false")
	}
}

func TestCopy(t *testing.T) {
	t.Parallel()

	nodeA := Node{Name: "A"}
	nodeB := Node{Name: "B"}
	nodeC := Node{Name: "C"}
	nodeD := Node{Name: "D"}

	edgeAB := Edge{
		Nodes:  [2]Node{nodeA, nodeB},
		Weight: 2,
	}

	edgeCD := Edge{
		Nodes:  [2]Node{nodeC, nodeD},
		Weight: 6,
	}

	graph := Graph{
		Nodes: []Node{nodeA, nodeB, nodeC, nodeD},
		Edges: []Edge{edgeAB, edgeCD},
	}

	graphCopy := graph.Copy()

	graph.Nodes[2], graph.Nodes[3] = graph.Nodes[3], graph.Nodes[2]

	if len(graphCopy.Nodes) != len(graph.Nodes) {
		t.Errorf("Copy did not work. Got %v instead of %v", len(graphCopy.Nodes), len(graph.Nodes))
	} else {
		if graphCopy.Nodes[0] != graph.Nodes[0] {
			t.Errorf("Copy did not work. Got %v instead of %v", graphCopy.Nodes[0], graph.Nodes[0])
		}

		if graphCopy.Nodes[1] != graph.Nodes[1] {
			t.Errorf("Copy did not work. Got %v instead of %v", graphCopy.Nodes[1], graph.Nodes[1])
		}

		if graphCopy.Nodes[2] != graph.Nodes[3] {
			t.Errorf("Copy did not work. Got %v instead of %v", graphCopy.Nodes[2], graph.Nodes[3])
		}

		if graphCopy.Nodes[3] != graph.Nodes[2] {
			t.Errorf("Copy did not work. Got %v instead of %v", graphCopy.Nodes[3], graph.Nodes[2])
		}
	}
}

func TestGeneratePathsWithoutEdgeRepetition(t *testing.T) {
	t.Parallel()

	nodeA := Node{Name: "A"}
	nodeB := Node{Name: "B"}
	nodeC := Node{Name: "C"}
	nodeD := Node{Name: "D"}
	nodeE := Node{Name: "E"}

	edgeAB := Edge{
		Nodes:  [2]Node{nodeA, nodeB},
		Weight: 5,
	}

	edgeBC := Edge{
		Nodes:  [2]Node{nodeB, nodeC},
		Weight: 4,
	}

	edgeCD := Edge{
		Nodes:  [2]Node{nodeC, nodeD},
		Weight: 8,
	}

	edgeDC := Edge{
		Nodes:  [2]Node{nodeD, nodeC},
		Weight: 8,
	}

	edgeDE := Edge{
		Nodes:  [2]Node{nodeD, nodeE},
		Weight: 6,
	}

	edgeAD := Edge{
		Nodes:  [2]Node{nodeA, nodeD},
		Weight: 5,
	}

	edgeCE := Edge{
		Nodes:  [2]Node{nodeC, nodeE},
		Weight: 2,
	}

	edgeEB := Edge{
		Nodes:  [2]Node{nodeE, nodeB},
		Weight: 3,
	}

	edgeAE := Edge{
		Nodes:  [2]Node{nodeA, nodeE},
		Weight: 7,
	}

	graph := Graph{
		Nodes: []Node{nodeA, nodeB, nodeC, nodeD, nodeE},
		Edges: []Edge{edgeAB, edgeBC, edgeCD, edgeDC, edgeDE, edgeAD, edgeCE, edgeEB, edgeAE},
	}

	// Case 1: no path found
	paths := graph.GeneratePathsWithoutEdgeRepetition(nodeA, nodeA)

	if len(paths) != 0 {
		t.Errorf("GeneratePathsWithoutEdgeRepetition did not work. Got %v instead of %v", len(paths), 0)
	}

	// Case 2: path found
	paths = graph.GeneratePathsWithoutEdgeRepetition(nodeC, nodeD)

	if len(paths) != 1 {
		t.Errorf("GeneratePathsWithoutEdgeRepetition did not work. Got %v instead of %v", len(paths), 1)
	} else {
		subgraph := paths[0].Subgraph

		if len(subgraph.Nodes) != 2 {
			t.Errorf("GeneratePathsWithoutEdgeRepetition did not work. Got %v instead of %v", len(subgraph.Nodes), 2)
		} else if subgraph.Nodes[0] != nodeC {
			t.Errorf("GeneratePathsWithoutEdgeRepetition did not work. Got %v instead of %v", subgraph.Nodes[0], nodeC)
		} else if subgraph.Nodes[1] != nodeD {
			t.Errorf("GeneratePathsWithoutEdgeRepetition did not work. Got %v instead of %v", subgraph.Nodes[1], nodeD)
		}

		if len(subgraph.Edges) != 1 {
			t.Errorf("GeneratePathsWithoutEdgeRepetition did not work. Got %v instead of %v", len(subgraph.Edges), 1)
		} else if !subgraph.Edges[0].Equals(&edgeCD) {
			t.Errorf("GeneratePathsWithoutEdgeRepetition did not work. Got %v instead of %v", subgraph.Edges[0], edgeCD)
		}
	}
}

func TestGeneratePathsWithMaxSteps(t *testing.T) {
	t.Parallel()

	nodeA := Node{Name: "A"}
	nodeB := Node{Name: "B"}
	nodeC := Node{Name: "C"}
	nodeD := Node{Name: "D"}
	nodeE := Node{Name: "E"}

	edgeAB := Edge{
		Nodes:  [2]Node{nodeA, nodeB},
		Weight: 5,
	}

	edgeBC := Edge{
		Nodes:  [2]Node{nodeB, nodeC},
		Weight: 4,
	}

	edgeCD := Edge{
		Nodes:  [2]Node{nodeC, nodeD},
		Weight: 8,
	}

	edgeDC := Edge{
		Nodes:  [2]Node{nodeD, nodeC},
		Weight: 8,
	}

	edgeDE := Edge{
		Nodes:  [2]Node{nodeD, nodeE},
		Weight: 6,
	}

	edgeAD := Edge{
		Nodes:  [2]Node{nodeA, nodeD},
		Weight: 5,
	}

	edgeCE := Edge{
		Nodes:  [2]Node{nodeC, nodeE},
		Weight: 2,
	}

	edgeEB := Edge{
		Nodes:  [2]Node{nodeE, nodeB},
		Weight: 3,
	}

	edgeAE := Edge{
		Nodes:  [2]Node{nodeA, nodeE},
		Weight: 7,
	}

	graph := Graph{
		Nodes: []Node{nodeA, nodeB, nodeC, nodeD, nodeE},
		Edges: []Edge{edgeAB, edgeBC, edgeCD, edgeDC, edgeDE, edgeAD, edgeCE, edgeEB, edgeAE},
	}

	// Case 1: path found in exact steps
	paths := graph.GeneratePathsWithMaxSteps(nodeA, nodeC, 3, true)

	if len(paths) != 1 {
		t.Errorf("GeneratePathsWithMaxSteps did not work. Got %v instead of %v", len(paths), 1)
	} else {
		subgraph := paths[0].Subgraph

		nodes := []Node{nodeA, nodeE, nodeB, nodeC}
		edges := []Edge{edgeAE, edgeEB, edgeBC}

		if len(subgraph.Nodes) != len(nodes) {
			t.Errorf("GeneratePathsWithMaxSteps did not work. Got %v instead of %v", len(subgraph.Nodes), len(nodes))
		} else {
			for i, n := range subgraph.Nodes {
				if n != nodes[i] {
					t.Errorf("GeneratePathsWithMaxSteps did not work. Got %v instead of %v", n, nodes[i])
				}
			}
		}

		if len(subgraph.Edges) != len(edges) {
			t.Errorf("GeneratePathsWithMaxSteps did not work. Got %v instead of %v", len(subgraph.Edges), len(edges))
		} else {
			for i, tr := range subgraph.Edges {
				if !tr.Equals(&edges[i]) {
					t.Errorf("GeneratePathsWithMaxSteps did not work. Got %v instead of %v", tr, edges[i])
				}
			}
		}
	}

	// Case 2: no path found in exact steps
	paths = graph.GeneratePathsWithMaxSteps(nodeA, nodeD, 2, true)

	if len(paths) != 0 {
		t.Errorf("GeneratePathsWithMaxSteps did not work. Got %v instead of %v", len(paths), 0)
	}

	// Case 3: path found in maximal steps
	paths = graph.GeneratePathsWithMaxSteps(nodeC, nodeE, 2, false)

	if len(paths) != 2 {
		t.Errorf("GeneratePathsWithMaxSteps did not work. Got %v instead of %v", len(paths), 2)
	} else {
		subgraphSlice := []Graph{paths[0].Subgraph, paths[1].Subgraph}

		if len(subgraphSlice[0].Nodes) > len(subgraphSlice[1].Nodes) {
			subgraphSlice[0], subgraphSlice[1] = subgraphSlice[1], subgraphSlice[0]
		}

		nodes := [][]Node{{nodeC, nodeE}, {nodeC, nodeD, nodeE}}
		edges := [][]Edge{{edgeCE}, {edgeCD, edgeDE}}

		for i, sg := range subgraphSlice {
			if len(sg.Nodes) != len(nodes[i]) {
				t.Errorf("GeneratePathsWithMaxSteps did not work. Got %v instead of %v", len(sg.Nodes), len(nodes[i]))
			} else {
				for j, n := range sg.Nodes {
					if n != nodes[i][j] {
						t.Errorf("GeneratePathsWithMaxSteps did not work. Got %v instead of %v", n, nodes[i][j])
					}
				}
			}

			if len(sg.Edges) != len(edges[i]) {
				t.Errorf("GeneratePathsWithMaxSteps did not work. Got %v instead of %v", len(sg.Edges), len(edges[i]))
			} else {
				for j, tr := range sg.Edges {
					if !tr.Equals(&edges[i][j]) {
						t.Errorf("GeneratePathsWithMaxSteps did not work. Got %v instead of %v", tr, edges[i][j])
					}
				}
			}
		}
	}

	// Case 5: no path found in maximal steps
	paths = graph.GeneratePathsWithMaxSteps(nodeA, nodeA, 2, false)

	if len(paths) != 0 {
		t.Errorf("GeneratePathsWithMaxSteps did not work. Got %v instead of %v", len(paths), 0)
	}
}

func TestGenerateSubgraphsWithMaxWeight(t *testing.T) {
	t.Parallel()

	nodeA := Node{Name: "A"}
	nodeB := Node{Name: "B"}
	nodeC := Node{Name: "C"}
	nodeD := Node{Name: "D"}
	nodeE := Node{Name: "E"}

	edgeAB := Edge{
		Nodes:  [2]Node{nodeA, nodeB},
		Weight: 5,
	}

	edgeBC := Edge{
		Nodes:  [2]Node{nodeB, nodeC},
		Weight: 4,
	}

	edgeCD := Edge{
		Nodes:  [2]Node{nodeC, nodeD},
		Weight: 8,
	}

	edgeDC := Edge{
		Nodes:  [2]Node{nodeD, nodeC},
		Weight: 8,
	}

	edgeDE := Edge{
		Nodes:  [2]Node{nodeD, nodeE},
		Weight: 6,
	}

	edgeAD := Edge{
		Nodes:  [2]Node{nodeA, nodeD},
		Weight: 5,
	}

	edgeCE := Edge{
		Nodes:  [2]Node{nodeC, nodeE},
		Weight: 2,
	}

	edgeEB := Edge{
		Nodes:  [2]Node{nodeE, nodeB},
		Weight: 3,
	}

	edgeAE := Edge{
		Nodes:  [2]Node{nodeA, nodeE},
		Weight: 7,
	}

	graph := Graph{
		Nodes: []Node{nodeA, nodeB, nodeC, nodeD, nodeE},
		Edges: []Edge{edgeAB, edgeBC, edgeCD, edgeDC, edgeDE, edgeAD, edgeCE, edgeEB, edgeAE},
	}

	// Case 1: no path found
	paths := graph.GeneratePathsWithMaxWeight(nodeA, nodeA, 30, true)

	if len(paths) != 0 {
		t.Errorf("GeneratePathsWithMaxWeight did not work. Got %v instead of %v", len(paths), 0)
	}

	// Case 2: path found
	paths = graph.GeneratePathsWithMaxWeight(nodeA, nodeD, 17, false)

	if len(paths) != 2 {
		t.Errorf("GeneratePathsWithMaxWeight did not work. Got %v instead of %v", len(paths), 2)
	} else {
		graphSlice := []Graph{paths[0].Subgraph, paths[1].Subgraph}

		if len(graphSlice[0].Nodes) > len(graphSlice[1].Nodes) {
			graphSlice[0], graphSlice[1] = graphSlice[1], graphSlice[0]
		}

		nodes := [][]Node{{nodeA, nodeD}, {nodeA, nodeB, nodeC, nodeD}}
		edges := [][]Edge{{edgeAD}, {edgeAB, edgeBC, edgeCD}}

		for i, gs := range graphSlice {
			if len(gs.Nodes) != len(nodes[i]) {
				t.Errorf("GeneratePathsWithMaxWeight did not work. Got %v instead of %v", len(gs.Nodes), len(nodes[i]))
			} else {
				for j, n := range gs.Nodes {
					if n != nodes[i][j] {
						t.Errorf("GeneratePathsWithMaxWeight did not work. Got %v instead of %v", n, nodes[i][j])
					}
				}
			}

			if len(gs.Edges) != len(edges[i]) {
				t.Errorf("GeneratePathsWithMaxWeight did not work. Got %v instead of %v", len(gs.Edges), len(edges[i]))
			} else {
				for j, tr := range gs.Edges {
					if !tr.Equals(&edges[i][j]) {
						t.Errorf("GeneratePathsWithMaxWeight did not work. Got %v instead of %v", tr, edges[i][j])
					}
				}
			}
		}
	}

	// Case 3: path found
	paths = graph.GeneratePathsWithMaxWeight(nodeA, nodeD, 17, true)

	if len(paths) != 1 {
		t.Errorf("GeneratePathsWithMaxWeight did not work. Got %v instead of %v", len(paths), 1)
	} else {
		subgraph := paths[0].Subgraph

		nodes := []Node{nodeA, nodeB, nodeC, nodeD}
		edges := []Edge{edgeAB, edgeBC, edgeCD}

		if len(subgraph.Nodes) != len(nodes) {
			t.Errorf("GeneratePathsWithMaxWeight did not work. Got %v instead of %v", len(subgraph.Nodes), len(nodes))
		} else {
			for i, n := range subgraph.Nodes {
				if n != nodes[i] {
					t.Errorf("GeneratePathsWithMaxWeight did not work. Got %v instead of %v", n, nodes[i])
				}
			}
		}

		if len(subgraph.Edges) != len(edges) {
			t.Errorf("GeneratePathsWithMaxWeight did not work. Got %v instead of %v", len(subgraph.Edges), len(edges))
		} else {
			for i, tr := range subgraph.Edges {
				if !tr.Equals(&edges[i]) {
					t.Errorf("GeneratePathsWithMaxWeight did not work. Got %v instead of %v", tr, edges[i])
				}
			}
		}
	}
}

func TestGenerateLowestHighestWeightPath(t *testing.T) {
	t.Parallel()

	nodeA := Node{Name: "A"}
	nodeB := Node{Name: "B"}
	nodeC := Node{Name: "C"}
	nodeD := Node{Name: "D"}
	nodeE := Node{Name: "E"}

	edgeAB := Edge{
		Nodes:  [2]Node{nodeA, nodeB},
		Weight: 5,
	}

	edgeBC := Edge{
		Nodes:  [2]Node{nodeB, nodeC},
		Weight: 4,
	}

	edgeCD := Edge{
		Nodes:  [2]Node{nodeC, nodeD},
		Weight: 8,
	}

	edgeDC := Edge{
		Nodes:  [2]Node{nodeD, nodeC},
		Weight: 8,
	}

	edgeDE := Edge{
		Nodes:  [2]Node{nodeD, nodeE},
		Weight: 6,
	}

	edgeAD := Edge{
		Nodes:  [2]Node{nodeA, nodeD},
		Weight: 5,
	}

	edgeCE := Edge{
		Nodes:  [2]Node{nodeC, nodeE},
		Weight: 2,
	}

	edgeEB := Edge{
		Nodes:  [2]Node{nodeE, nodeB},
		Weight: 3,
	}

	edgeAE := Edge{
		Nodes:  [2]Node{nodeA, nodeE},
		Weight: 7,
	}

	graph := Graph{
		Nodes: []Node{nodeA, nodeB, nodeC, nodeD, nodeE},
		Edges: []Edge{edgeAB, edgeBC, edgeCD, edgeDC, edgeDE, edgeAD, edgeCE, edgeEB, edgeAE},
	}

	// Case 1: HCP
	paths := graph.GenerateLowestHighestWeightPath(nodeB, nodeB, true)

	if paths[0].Weight != 9 {
		t.Errorf("GenerateLowestHighestWeightPath did not work. Got %v instead of %v", paths[0].Weight, 9)
	}

	// Case 2: TSP
	paths = graph.GenerateLowestHighestWeightPath(nodeA, nodeD, false)

	if paths[0].Weight != 22 {
		t.Errorf("GenerateLowestHighestWeightPath did not work. Got %v instead of %v", paths[0].Weight, 22)
	}
}

func TestGenerateShortestLongestPath(t *testing.T) {
	t.Parallel()

	nodeA := Node{Name: "A"}
	nodeB := Node{Name: "B"}
	nodeC := Node{Name: "C"}
	nodeD := Node{Name: "D"}
	nodeE := Node{Name: "E"}

	edgeAB := Edge{
		Nodes:  [2]Node{nodeA, nodeB},
		Weight: 5,
	}

	edgeBC := Edge{
		Nodes:  [2]Node{nodeB, nodeC},
		Weight: 4,
	}

	edgeCD := Edge{
		Nodes:  [2]Node{nodeC, nodeD},
		Weight: 8,
	}

	edgeDC := Edge{
		Nodes:  [2]Node{nodeD, nodeC},
		Weight: 8,
	}

	edgeDE := Edge{
		Nodes:  [2]Node{nodeD, nodeE},
		Weight: 6,
	}

	edgeAD := Edge{
		Nodes:  [2]Node{nodeA, nodeD},
		Weight: 5,
	}

	edgeCE := Edge{
		Nodes:  [2]Node{nodeC, nodeE},
		Weight: 2,
	}

	edgeEB := Edge{
		Nodes:  [2]Node{nodeE, nodeB},
		Weight: 3,
	}

	edgeAE := Edge{
		Nodes:  [2]Node{nodeA, nodeE},
		Weight: 7,
	}

	graph := Graph{
		Nodes: []Node{nodeA, nodeB, nodeC, nodeD, nodeE},
		Edges: []Edge{edgeAB, edgeBC, edgeCD, edgeDC, edgeDE, edgeAD, edgeCE, edgeEB, edgeAE},
	}

	paths := graph.GenerateShortestLongestPath(nodeB, nodeB, true)

	if len(paths[0].Subgraph.Edges) != 3 {
		t.Errorf("GenerateShortestLongestPath did not work. Got %v instead of %v", len(paths[0].Subgraph.Edges), 3)
	}

	paths = graph.GenerateShortestLongestPath(nodeA, nodeD, false)

	if len(paths[0].Subgraph.Edges) != 4 {
		t.Errorf("GenerateShortestLongestPath did not work. Got %v instead of %v", len(paths[0].Subgraph.Edges), 4)
	}
}
