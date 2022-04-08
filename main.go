package main

import (
	"fmt"
	"thesisGoRemake/FABuilder/Builder"
	"thesisGoRemake/FABuilder/Builder/CharSet"
	"thesisGoRemake/FABuilder/DOT"
)

func main() {
	var d Builder.DFABuilderListener

	d.CurrGraph = Builder.NewGraph("", "")
	graph := d.CurrGraph

	// /aab/
	//
	//	1-a-2-a-3-b-[4]
	//
	//

	v1 := graph.CreateDefaultVertex()
	v2 := graph.CreateDefaultVertex()
	v3 := graph.CreateDefaultVertex()
	v4 := graph.CreateDefaultVertex()
	v5 := graph.CreateTerminalVertex()

	graph.RootVertex = v1

	e12 := graph.CreateEdge(v1, v2, CharSet.NewCharSet(CharSet.CharInterval{Start: 'a', Stop: 'a'}))
	graph.CreateEdge(v2, v3, CharSet.NewCharSet(CharSet.CharInterval{Start: 'a', Stop: 'a'}))
	graph.CreateEdge(v3, v4, CharSet.NewCharSet(CharSet.CharInterval{Start: 'a', Stop: 'a'}))
	graph.CreateEdge(v4, v5, CharSet.NewCharSet(CharSet.CharInterval{Start: 'b', Stop: 'b'}))

	//loopback edge
	graph.CopyEdge(v2, v2, *e12)

	DOT.Export("./before.DOT", "a", *graph)

	d.EnforceDeterminism()

	//Expected form:
	//         -a-
	//         \ /
	//	1-a-2-a-3-b-[4]
	//	    ^	     |
	//      |        |
	//      ----a----

	DOT.Export("./after.DOT", "a", *graph)

	var visitor Builder.GraphWalker

	err := visitor.Walk(*graph, "aaaab")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	if len(visitor.GetMatches()) == 0 {
		fmt.Printf("No matches found, visitor did not reach vertex")
		return
	}
}
