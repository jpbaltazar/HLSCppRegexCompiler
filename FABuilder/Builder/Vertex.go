package Builder

import "fmt"

type VertexID uint64

type Vertex struct {
	Id VertexID

	//makes it efficient to add/remove the edges
	Outgoing map[VertexID]*Edge // SOURCE -> EDGE
	Incoming map[VertexID]*Edge // TARGET -> EDGE

	Terminal     bool
	EndOfSubject bool
}

func (v Vertex) GetIncoming() []*Edge {
	var edges []*Edge

	for _, e := range v.Incoming {
		edges = append(edges, e)
	}

	return edges
}

func (v Vertex) GetOutgoing() []*Edge {
	var edges []*Edge

	for _, e := range v.Outgoing {
		edges = append(edges, e)
	}

	return edges
}

func (v Vertex) DebugString() string {
	/*str := fmt.Sprintf("V%d[%d:%d]", v.Id, len(v.Incoming), len(v.Outgoing))
	if v.Terminal{
		str += "T"
	}
	if v.EndOfSubject{
		str += "E"
	}

	return fmt.Sprintf("%s", str)//str */

	return fmt.Sprintf("V%d", v.Id)
}
