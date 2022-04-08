package Builder

import (
	"testing"
	"thesisGoRemake/FABuilder/Builder/CharSet"
	"thesisGoRemake/FABuilder/Flags"
)

func TestNewGraph(t *testing.T) {
	g := NewGraphWithFlags("test", Flags.NewFlags(""))

	if g.Rule != "test" {
		t.Errorf("Failed to set Rule, expected %s, got %s\n", "test", g.Rule)
	}

	var emptyFlags Flags.Flags
	if g.Flags != emptyFlags {
		t.Errorf("Flags weren't empty\n")
	}

	if len(g.vertices) != 0 {
		t.Errorf("Already has vertices\n")
	}

	if len(g.Edges()) != 0 {
		t.Errorf("Already has edges\n")
	}

	if g.RootVertex != nil {
		t.Errorf("RootVertex is nil\n")
	}

	if g.ErrorVertex != nil {
		t.Errorf("ErrorVertex is nil\n")
	}

}

func TestGraph_CreateVertex(t *testing.T) {
	g := NewGraphWithFlags("", Flags.NewFlags(""))

	v := g.CreateVertex(false, false)
	if len(g.vertices) != 1 {
		t.Errorf("Failed to create vertex\n")
	}

	if len(g.GetVertexMap()) != 1 {
		t.Errorf("GetVertexList() call failed to return length 1")
	}

	if v.Terminal != false || v.EndOfSubject != false {
		t.Errorf("Vertex created with wrong properties\n")
	}

	if len(v.Outgoing) != 0 || len(v.Incoming) != 0 {
		t.Errorf("Incoming/Outgoing maps are not empty\n")
	}
}

func TestGraph_CreateEdge(t *testing.T) {
	g := NewGraphWithFlags("", Flags.NewFlags(""))

	v1 := g.CreateVertex(false, false)
	v2 := g.CreateVertex(false, false)

	e := g.CreateEdge(v1, v2, CharSet.NewEmptyCharSet())
	v1.Outgoing[v2.Id] = e
	v2.Incoming[v1.Id] = e

	//
	if len(g.Edges()) != 1 {
		t.Errorf("Failed to create edge\n")
	}

	if e.From.Id != v1.Id {
		t.Errorf("Wrong vertex set as origin vertex\n")
	}

	if e.To.Id != v2.Id {
		t.Errorf("Wrong vertex set as destination vertex\n")
	}
}

func TestGraph_RemoveEdge(t *testing.T) {
	g := NewGraphWithFlags("", Flags.NewFlags(""))

	v1 := g.CreateVertex(false, false)
	v2 := g.CreateVertex(false, false)

	e := g.CreateEdge(v1, v2, &CharSet.CharSet{})

	g.RemoveEdge(*e)

	if len(v1.Outgoing) != 0 {
		t.Errorf("Failed to remove reference in v1")
	}

	if len(v2.Incoming) != 0 {
		t.Errorf("Failed to remove reference in v2")
	}
}

func TestGraph_RemoveVertex(t *testing.T) {
	g := NewGraphWithFlags("", Flags.NewFlags(""))

	v1 := g.CreateVertex(false, false)
	v2 := g.CreateVertex(false, false)

	g.CreateEdge(v1, v2, &CharSet.CharSet{})

	g.RemoveVertex(v2)

	if len(g.vertices) != 1 {
		t.Errorf("Failed to remove reference in vertex map")
	}

	if len(v1.Outgoing) != 0 {
		t.Errorf("Failed to remove reference in v1")
	}

	if len(g.Edges()) != 0 {
		t.Errorf("Failed to remove edges")
	}

}
