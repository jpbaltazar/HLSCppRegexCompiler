package DOT

import (
	"fmt"
	"os"
	"text/template"
	"thesisGoRemake/FABuilder/Builder"
	Parsers2 "thesisGoRemake/FABuilder/Builder/CharSet/Parsers"
)

type exporter struct {
	g      Builder.Graph
	parser Parsers2.Parser
}

func (exp *exporter) PrintEdge(e Builder.Edge) string {
	label := ""
	if !e.IsBackreference {
		str := ""
		suffix := ""

		//charset
		if len(e.C.Intervals) != 1 {
			str = "["
			suffix = "]"
		}
		for _, i := range e.C.Intervals {
			str += exp.parser.IntervalToString(i)
		}

		label += str + suffix

		if !e.LoopInteractions.IsEmpty() {
			for _, i := range e.LoopInteractions.GetInteractions() {
				label += i.ToString() + " "
			}

			label += "\n"
		}

		if !e.LoopConditions.IsEmpty() {
			for _, c := range e.LoopConditions.GetConditions() {
				label += c.ToString() + " "
			}

			label += "\n"
		}

	} else {
		label += "Backref: (" + e.BackrefRef + ")"
	}

	//capture groups
	if len(e.CaptureGroups) > 0 {
		label += "("

		for _, c := range e.CaptureGroups {
			label += c + ","
		}

		label += ")\n"
	}

	return fmt.Sprintf("%d -> %d [label=\"%s\"];", e.From.Id, e.To.Id, label)
}

func (exp *exporter) PrintVertex(v Builder.Vertex) string {
	str := fmt.Sprintf("%d [label=\"%s\"", v.Id, v.DebugString())
	if v.Terminal {
		str += ", peripheries=2"
	}
	str += "];"

	return str
}

func Export(FilePath, name string, g Builder.Graph) {
	files, err := template.ParseFiles("FABuilder/DOT/DOTTemplate.txt")
	if err != nil {
		return
	}

	exp := exporter{
		g:      g,
		parser: Parsers2.ASCII{},
	}

	verticesEdges := ""

	for _, v := range g.GetVertexMap() {
		verticesEdges += "\t" + exp.PrintVertex(*v) + "\n"
	}

	for _, e := range g.Edges() {
		verticesEdges += "\t" + exp.PrintEdge(*e) + "\n"
	}

	t := struct {
		Name          string
		VerticesEdges string
	}{
		name,
		verticesEdges,
	}

	f, err := os.Create(FilePath)
	if err != nil {
		return
	}

	err = files.Execute(f, t)
	if err != nil {
		return
	}
}
