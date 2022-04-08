package Builder

import (
	"thesisGoRemake/FABuilder/Builder/CharSet"
	Flags "thesisGoRemake/FABuilder/Flags"
	"thesisGoRemake/FABuilder/NumberedLoops"
)

type Graph struct {
	//definitions
	RootVertex  *Vertex
	ErrorVertex *Vertex

	Rule  string
	Flags Flags.Flags

	//TODO shorthand list of references used (for backreference listing)

	//private internals
	vertices map[VertexID]*Vertex

	variables map[NumberedLoops.VarId]*NumberedLoops.LoopVar

	CurrGivenVertexId  VertexID            //lazy id'ing of vertices, increment one by each new, any graph won't have that many
	CurrGivenLoopVarId NumberedLoops.VarId //lazy id'ing of Loop variables
}

//New Graphs

func NewGraph(rule string, flags string) *Graph {
	return NewGraphWithFlags(rule, Flags.NewFlags(flags))
}

func NewGraphWithFlags(rule string, flags Flags.Flags) *Graph {
	g := new(Graph)
	*g = Graph{
		nil,
		nil,
		rule,
		flags,
		make(map[VertexID]*Vertex),
		make(map[NumberedLoops.VarId]*NumberedLoops.LoopVar),
		0,
		0,
	}

	return g
}

//Vertex Methods

func (g *Graph) CreateDefaultVertex() *Vertex {
	return g.CreateVertex(false, false)
}

func (g *Graph) CreateTerminalVertex() *Vertex {
	return g.CreateVertex(true, false)
}

func (g *Graph) CreateVertex(terminal, endOfSubject bool) *Vertex {
	v := new(Vertex)
	*v = Vertex{
		Id:           VertexID(g.CurrGivenVertexId),
		Incoming:     make(map[VertexID]*Edge),
		Outgoing:     make(map[VertexID]*Edge),
		Terminal:     terminal,
		EndOfSubject: endOfSubject,
	}

	g.vertices[v.Id] = v

	//next
	g.CurrGivenVertexId++

	return v
}

func (g *Graph) RemoveVertex(vertex *Vertex) {
	for _, edge := range vertex.Incoming {
		if _, exists := edge.From.Outgoing[vertex.Id]; exists {
			delete(edge.From.Outgoing, edge.To.Id)
		}
	}

	for _, edge := range vertex.Outgoing {
		if _, exists := edge.To.Incoming[vertex.Id]; exists {
			delete(edge.To.Incoming, edge.From.Id)
		}
	}

	delete(g.vertices, vertex.Id)
}

func (g *Graph) RemoveVertexFromID(id VertexID) {
	g.RemoveVertex(g.vertices[id])
}

func (g *Graph) MergeVertexBIntoA(A, B *Vertex) {
	g.RemoveVertex(g.CopyVertexBIntoA(A, B))
}

func (g *Graph) CopyVertexBIntoA(A, B *Vertex) *Vertex {
	A.Terminal = A.Terminal && B.Terminal

	for _, edge := range B.Outgoing {
		g.CopyEdge(A, edge.To, *edge)
	}

	for _, edge := range B.Incoming {
		existingEdge := A.Incoming[edge.To.Id]

		if existingEdge != nil { //already exists
			if edge.C.EqualTo(*existingEdge.C) {

				existingEdge.CaptureGroups =
					append(existingEdge.CaptureGroups, edge.CaptureGroups...)
				existingEdge.LoopInteractions.AddInteractions(*edge.LoopInteractions)
			}
		}

		g.CopyEdge(edge.From, B, *edge)
	}

	return B
}

func (g Graph) GetVertexMap() map[VertexID]*Vertex {
	return g.vertices
}

func (g Graph) GetVertexList() []*Vertex {
	//keys := reflect.ValueOf(g.vertices).MapKeys()
	vertices := make([]*Vertex, len(g.vertices))
	i := 0
	for _, val := range g.vertices {
		vertices[i] = val
		i++
	}
	return vertices
}

func (g Graph) GetVertex(id VertexID) *Vertex {
	return g.vertices[id]
}

//Edge Methods

func (g *Graph) CreateBackrefEdge(from, to *Vertex, ref string) *Edge {
	e := &Edge{
		From:             from,
		To:               to,
		IsBackreference:  true,
		BackrefRef:       ref,
		LoopConditions:   NumberedLoops.NewLoopConditions(),
		LoopInteractions: NumberedLoops.NewLoopInteractions(),
	}

	from.Outgoing[to.Id] = e
	to.Incoming[from.Id] = e
	return e
}

func (g *Graph) CreateEdge(from, to *Vertex, conditions *CharSet.CharSet) *Edge {
	e := &Edge{
		From:             from,
		To:               to,
		C:                conditions,
		LoopConditions:   NumberedLoops.NewLoopConditions(),
		LoopInteractions: NumberedLoops.NewLoopInteractions(),
	}

	from.Outgoing[to.Id] = e
	to.Incoming[from.Id] = e

	return e
}

func (g *Graph) CreateEpsilonEdge(from, to *Vertex) *Edge {
	return g.CreateEdge(from, to, &CharSet.CharSet{Intervals: []CharSet.CharInterval{}})
}

func (g *Graph) CopyEdge(from, to *Vertex, edge Edge) *Edge {
	e := Edge{
		From:             from,
		To:               to,
		IsBackreference:  edge.IsBackreference,
		BackrefRef:       edge.BackrefRef,
		CaptureGroups:    append([]string{}, edge.CaptureGroups...),
		C:                CharSet.CopyCharSet(*edge.C),
		LoopConditions:   NumberedLoops.CopyLoopConditions(*edge.LoopConditions),
		LoopInteractions: NumberedLoops.CopyLoopInteractions(*edge.LoopInteractions),
	}

	from.Outgoing[to.Id] = &e
	to.Incoming[from.Id] = &e

	return &e
}

func (g *Graph) RemoveEdge(e Edge) {
	delete(e.From.Outgoing, e.To.Id)
	delete(e.To.Incoming, e.From.Id)
}

func (g Graph) Edges() []*Edge {
	var edges []*Edge

	for _, v := range g.GetVertexMap() {
		for _, e := range v.GetOutgoing() {
			edges = append(edges, e)
		}
	}

	return edges
}
