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

	for _, char := range input {
		pathFound := false

		for _, out := range currVertex.Outgoing {
			if out.C.Contains(uint64(char)) && out.LoopConditions.Evaluate(visitor.varTable) {
				pathFound = true
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

		if !pathFound {
			return fmt.Errorf("no path found, from given input")
		}
	}

	return nil
}

func (visitor GraphWalker) GetMatches() []Match {
	return append([]Match{}, visitor.matches...)
}
