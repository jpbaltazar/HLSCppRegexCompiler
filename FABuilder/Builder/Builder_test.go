package Builder

import (
	"fmt"
	"testing"
	"thesisGoRemake/FABuilder"
	"thesisGoRemake/FABuilder/Builder/CharSet"
	"thesisGoRemake/FABuilder/Builder/CharSet/Parsers"
	"thesisGoRemake/FABuilder/Flags"
)

func TestSuperSetSolver_basic(t *testing.T) {
	g := NewGraphWithFlags("", Flags.NewFlags(""))
	d := &DFABuilderListener{
		Graphs:            make([]Graph, 0),
		CurrGraph:         nil,
		Parser:            Parsers.ASCII{}, //default
		SubGraphStack:     *NewSubGraphStack(),
		CaptureGroupStack: FABuilder.NewCaptureGroupStack(),
	}
	d.CurrGraph = g

	v1 := g.CreateDefaultVertex()
	v2 := g.CreateDefaultVertex()
	v3 := g.CreateTerminalVertex()

	g.CreateEpsilonEdge(v1, v2)
	g.CreateEdge(v2, v3, &CharSet.CharSet{
		Intervals: []CharSet.CharInterval{
			{'a', 'a'},
		},
	})

	//1-epsilon->2-->3
	d.SuperSetSolver(v1)

	vertices := d.CurrGraph.GetVertexMap()

	if len(vertices) != 2 {
		t.Errorf("Expected 2 vertices, got %d\n", len(vertices))
	}

	var start *Vertex
	var stop *Vertex

	start = nil
	stop = nil
	for _, v := range vertices {
		if v == nil {
			t.Errorf("Vertex was nil")
			return
		}

		if v.Terminal {
			stop = v
		} else {
			start = v
		}
	}

	if start == nil {
		t.Errorf("Start Vertex not found")
		return
	}

	if stop == nil {
		t.Errorf("Stop Vertex not found")
		return
	}

	if len(start.GetIncoming()) != 0 {
		t.Errorf("Expected no incoming edges for the start vertex, got %d", len(start.GetIncoming()))
		return
	}

	if len(start.GetOutgoing()) != 1 {
		t.Errorf("Expected 1 outgoing edge for the start vertex, got %d", len(start.GetOutgoing()))
		return
	}

	if len(stop.GetIncoming()) != 1 {
		t.Errorf("Expected 1 incoming edge for the stop vertex, got %d", len(stop.GetIncoming()))
		return
	}

	if len(stop.GetOutgoing()) != 0 {
		t.Errorf("Expected no outgoing edges for the stop vertex, got %d", len(stop.GetOutgoing()))
		return
	}

	edges := d.CurrGraph.Edges()
	if len(edges) != 1 {
		t.Errorf("Expected 1 edge, got %d\n", len(edges))
		return
	}

	if edges[0].C.IsEmpty() {
		t.Errorf("Expected 'a' edge, got epsilon?")
		return
	}
}

func TestEnforceDeterminism_basic(t *testing.T) {
	var d DFABuilderListener

	d.CurrGraph = NewGraph("", "")
	graph := d.CurrGraph

	// Original
	//         v3
	//       /a  \c
	//v1-a-v2      [v5]
	//       \b  /d
	//         v4

	v1 := graph.CreateDefaultVertex()
	v2 := graph.CreateDefaultVertex()
	v3 := graph.CreateDefaultVertex()
	v4 := graph.CreateDefaultVertex()
	v5 := graph.CreateDefaultVertex()

	graph.RootVertex = v1

	graph.CreateEdge(v1, v2, CharSet.NewCharSet(CharSet.CharInterval{Start: 'a', Stop: 'a'}))
	graph.CreateEdge(v2, v3, CharSet.NewCharSet(CharSet.CharInterval{Start: 'a', Stop: 'a'}))
	graph.CreateEdge(v2, v4, CharSet.NewCharSet(CharSet.CharInterval{Start: 'b', Stop: 'b'}))
	graph.CreateEdge(v3, v5, CharSet.NewCharSet(CharSet.CharInterval{Start: 'c', Stop: 'c'}))
	graph.CreateEdge(v4, v5, CharSet.NewCharSet(CharSet.CharInterval{Start: 'd', Stop: 'd'}))

	d.EnforceDeterminism()

	// Expected
	//         v3
	//       /a  \c
	//v1-a-v2      [v5]
	//       \b  /d
	//         v4

	if len(graph.GetVertexMap()) != 5 {
		t.Errorf("Wrong number of resulting vertices\n")
		return
	}

	var newV1, newV2, newV3, newV4, newV5 *Vertex

	visited := make(map[VertexID]bool)

	VStack := NewVertexStack()
	VStack.Push(graph.RootVertex)
	for VStack.Size() > 0 {
		V := VStack.Pop()

		switch len(V.Incoming) {
		case 0:
			newV1 = V

		case 1:
			switch len(V.Outgoing) {
			case 1:
				for _, edge := range V.Outgoing {
					if edge.C.EqualTo(CharSet.CharSet{Intervals: []CharSet.CharInterval{{'d', 'd'}}}) {
						newV4 = V
					} else {
						newV3 = V
					}
				}
			case 2:
				newV2 = V
			}
		case 2:
			newV5 = V
		}

		for id := range V.Outgoing {
			if _, ok := visited[id]; !ok {
				VStack.Push(graph.GetVertex(id))
				visited[id] = true
			}
		}
	}

	if newV1 == nil ||
		newV2 == nil ||
		newV3 == nil ||
		newV4 == nil ||
		newV5 == nil {

		t.Errorf("Unexpected graph construction")
		return
	}

}

func TestEnforceDeterminism_InternalMatching(t *testing.T) {
	var d DFABuilderListener

	d.CurrGraph = NewGraph("", "")
	graph := d.CurrGraph

	// /aab/
	//
	//	1-a-2-a-3-b-[4]
	//

	v1 := graph.CreateDefaultVertex()
	v2 := graph.CreateDefaultVertex()
	v3 := graph.CreateDefaultVertex()
	v4 := graph.CreateTerminalVertex()

	graph.RootVertex = v1

	e12 := graph.CreateEdge(v1, v2, CharSet.NewCharSet(CharSet.CharInterval{Start: 'a', Stop: 'a'}))
	graph.CreateEdge(v2, v3, CharSet.NewCharSet(CharSet.CharInterval{Start: 'a', Stop: 'a'}))
	graph.CreateEdge(v3, v4, CharSet.NewCharSet(CharSet.CharInterval{Start: 'b', Stop: 'b'}))

	//loopback edge
	graph.CopyEdge(v2, v2, *e12)

	d.EnforceDeterminism()

	//Expected form:
	//         -a-
	//         \ /
	//	1-a-2-a-3-b-[4]

	var walker GraphWalker
	err := walker.Walk(*graph, "aaab")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	if len(walker.GetMatches()) != 1 {
		fmt.Printf("Expected one match")
		return
	}

}
