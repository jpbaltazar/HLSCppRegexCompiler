package Builder

import (
	"fmt"
	"thesisGoRemake/FABuilder/NumberedLoops"
)

type Match struct {
	Vertex Vertex
	Path   string
}

type GraphWalker struct {
	varTable map[NumberedLoops.VarId]int64

	matches []Match

	valid bool
	error string
}

//TODO implement backreferences
func (visitor *GraphWalker) Walk(graph Graph, input string) error {
	currVertex := graph.RootVertex

	currInput := ""

	referenceBank := make(map[string]string)

	for i := 0; i < len(input); i++ {
		char := input[i]
		var takenEdge *Edge = nil

		for _, out := range currVertex.Outgoing {
			if out.IsBackreference {

				for j := 0; j < len(referenceBank[out.BackrefRef]); j++ {
					c := input[i]

					if string(c) == string(referenceBank[out.BackrefRef][j]) {
						referenceBank[out.BackrefRef] = referenceBank[out.BackrefRef][1:]
					} else {
						return fmt.Errorf("backreference did not match input")
					}

					i++
				}

				if string(char) == string(referenceBank[out.BackrefRef][0]) {
					referenceBank[out.BackrefRef] = referenceBank[out.BackrefRef][1:]
				} else {
					return fmt.Errorf("backreference did not match input")
				}
			} else {
				if out.C.Contains(uint64(char)) && out.LoopConditions.Evaluate(visitor.varTable) {
					takenEdge = out
					currInput += string(char)

					currVertex = out.To
					if currVertex.Terminal {
						visitor.matches = append(visitor.matches, Match{
							*out.To,
							currInput,
						})
					}

					for _, interaction := range out.LoopInteractions.GetInteractions() {
						if interaction.Type == NumberedLoops.AddToVar {
							visitor.varTable[interaction.Var.Id] += interaction.Value
						} else { //set
							visitor.varTable[interaction.Var.Id] = interaction.Value
						}
					}

					break
				}
			}
		}

		if takenEdge == nil { //no edge taken
			return fmt.Errorf("no path found, from given input")
		} else { //add to reference banks if useful
			for _, reference := range takenEdge.CaptureGroups {
				if _, ok := graph.referencesUsed[reference]; ok {
					referenceBank[reference] += string(char)
				}
			}
		}
	}

	return nil
}

func (visitor GraphWalker) GetMatches() []Match {
	return append([]Match{}, visitor.matches...)
}
