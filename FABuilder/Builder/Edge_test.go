package Builder

import (
	"testing"
	"thesisGoRemake/FABuilder/Flags"
)

func TestEdge_IsEpsilon(t *testing.T) {
	g := NewGraphWithFlags("", Flags.NewFlags(""))

	v1 := g.CreateDefaultVertex()
	v2 := g.CreateTerminalVertex()

	e := g.CreateEpsilonEdge(v1, v2)

	if !e.IsEpsilon() {
		t.Errorf("Failed to create an epsilon edge")
	}
}
