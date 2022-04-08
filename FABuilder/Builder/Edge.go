package Builder

import "C"
import (
	"thesisGoRemake/FABuilder/Builder/CharSet"
	"thesisGoRemake/FABuilder/NumberedLoops"
)

type Edge struct {
	From *Vertex
	To   *Vertex

	IsBackreference bool
	BackrefRef      string

	CaptureGroups []string

	C *CharSet.CharSet //e.g. [['a':'d'], ['A':'Z'], ['_':'_']

	LoopConditions   *NumberedLoops.LoopConditions   //e.g. i < 20
	LoopInteractions *NumberedLoops.LoopInteractions //e.g. i += 1
}

func (e Edge) IsEpsilon() bool {
	return e.C.IsEmpty() && e.LoopInteractions.IsEmpty()
}

func (e Edge) IsEpsilonWithInteractions() bool {
	return e.C.IsEmpty() && !e.LoopInteractions.IsEmpty()
}
