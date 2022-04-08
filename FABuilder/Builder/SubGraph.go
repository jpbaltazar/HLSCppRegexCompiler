package Builder

import (
	"thesisGoRemake/FABuilder/NumberedLoops"
)

type SubGraphType int

const (
	Alternation SubGraphType = iota
	Expr
	Atom
	Cc_atom
	Literal
	SEPARATOR
)

type SubGraph struct {
	Start, Stop *Vertex
	Type        SubGraphType
	Ctx         string
}

func (g *SubGraph) Decompose() (*Vertex, *Vertex) {
	return g.Start, g.Stop
}

func getVertexOfVertex(g Graph, vertexTranslationTable map[VertexID]VertexID, vertexID VertexID) *Vertex {
	if _, ok := vertexTranslationTable[vertexID]; !ok {
		originalVertex := g.GetVertex(vertexID)
		newVertex := g.CreateVertex(originalVertex.Terminal, originalVertex.EndOfSubject)

		//we don't copy the incoming/outgoing edges as they will be filed in the Copy Subgraph function

		vertexTranslationTable[vertexID] = newVertex.Id
	}

	return g.GetVertex(vertexTranslationTable[vertexID])
}

func getVarOfVar(g Graph, varTranslationTable map[*NumberedLoops.LoopVar]*NumberedLoops.LoopVar, Variable *NumberedLoops.LoopVar) *NumberedLoops.LoopVar {
	if _, ok := varTranslationTable[Variable]; !ok {
		varTranslationTable[Variable] = &NumberedLoops.LoopVar{
			Id:            g.CurrGivenLoopVarId,
			CaptureGroups: append([]string{}, Variable.CaptureGroups...),
			Min:           Variable.Min,
			Max:           Variable.Max,
			Type:          Variable.Type,
		}

		g.CurrGivenLoopVarId++
	}

	return varTranslationTable[Variable]
}

//	useful for (---){x,} quantifier, simply turn it into a (---){x}(---)*
//	removes complex conditions
func (s *SubGraph) Copy(graph Graph) *SubGraph {
	//Deep copy of subgraph
	//Make new vertices
	//Use new variables

	vertexToVertex := make(map[VertexID]VertexID)
	//varToVar := make(map[*NumberedLoops.LoopVar]*NumberedLoops.LoopVar)

	originalRoot := s.Start
	newRoot := getVertexOfVertex(graph, vertexToVertex, originalRoot.Id)

	vertexStack := NewVertexStack()
	vertexStack.Push(newRoot)
	for vertexStack.Size() > 0 {
		v := vertexStack.Pop()

		for _, out := range v.Outgoing {
			target := getVertexOfVertex(graph, vertexToVertex, out.To.Id)

			edge := Edge{
				From:             v,
				To:               target,
				IsBackreference:  out.IsBackreference,
				BackrefRef:       out.BackrefRef,
				CaptureGroups:    append([]string{}, out.CaptureGroups...),
				C:                out.C.Copy(),
				LoopConditions:   nil,
				LoopInteractions: nil,
			}

			//TODO fix variables
			//	Maybe create a variable table for graph
			//	Also add loop conditions/interactions

			/*for _, a := range out.LoopConditions.GetConditions() {
				edge.LoopConditions.Add(getVarOfVar(graph, varToVar, a.Var))
			}*/

			graph.CopyEdge(v, target, edge)
		}
	}

	subgraph := &SubGraph{
		Start: getVertexOfVertex(graph, vertexToVertex, s.Start.Id),
		Stop:  getVertexOfVertex(graph, vertexToVertex, s.Stop.Id),
		Type:  s.Type,
		Ctx:   s.Ctx,
	}

	return subgraph
}
