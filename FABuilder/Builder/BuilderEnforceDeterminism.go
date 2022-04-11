package Builder

import "thesisGoRemake/FABuilder/Builder/CharSet"

//Essentially looks for edges that have the same characters and lead to different vertices and creates middle vertices
//that point to both vertices with the intersection of the edges

type CharSetPartition struct {
	charset CharSet.CharSet
	edges   []*Edge
}

func (d *DFABuilderListener) EnforceDeterminism() error {
	//stuff for translating
	//[Vertex] -> SuperVertex -> Corresponding Vertex
	SVLookup := NewSuperVertexLookup()
	namingSVId := SuperVertexID(0)

	//initialize the visitation stack
	rootSV := d.getSuperVertexOfVertexSet(SVLookup, &namingSVId, []*Vertex{d.CurrGraph.RootVertex})

	SVStack := NewSuperVertexStack()
	SVStack.Push(rootSV)

	//keep a list of the original vertices for deletion
	originalVertices := make(map[VertexID]bool)
	originalVertices[d.CurrGraph.RootVertex.Id] = true

	//replace the root vertex
	superVertexToVertex := make(map[SuperVertexID]*Vertex)
	newRoot := d.getVertexOfSuperVertex(superVertexToVertex, *rootSV)
	d.CurrGraph.RootVertex = newRoot

	visitedSuperVertices := make(map[SuperVertexID]bool)

	for SVStack.Size() > 0 {
		SV := SVStack.Pop()

		charsetPartitions := make([]*CharSetPartition, 0)

		//Partition the charset
		for _, v := range SV.vertices {
			for _, out := range v.Outgoing {
				originalVertices[out.To.Id] = true

				outCSP := &CharSetPartition{
					charset: *out.C,
					edges:   []*Edge{out},
				}

				found := false
				for _, p := range charsetPartitions {
					newP, newOut, intersection := p.charset.Intersect(outCSP.charset)

					if !intersection.IsEmpty() {
						found = true

						//Change P partition
						*p = CharSetPartition{
							charset: newP,
							edges:   p.edges,
						}

						//Change Out Partition
						outCSP.charset = newOut

						//Add intersection
						charsetPartitions = append(charsetPartitions, &CharSetPartition{
							intersection,
							append(append([]*Edge{}, p.edges...), outCSP.edges...),
						})
					}
				}

				//directly add partition
				if !found {
					charsetPartitions = append(charsetPartitions, outCSP)
				}
			}
		}

		//remove empty partitions
		cleanPartitions := make([]*CharSetPartition, 0)
		for _, p := range charsetPartitions {
			if !p.charset.IsEmpty() {
				cleanPartitions = append(cleanPartitions, p)
			}
		}
		charsetPartitions = cleanPartitions

		//Assemble the corresponding transitions
		for _, p := range charsetPartitions {
			vertexSet := make([]*Vertex, 0)
			for _, e := range p.edges {
				vertexSet = append(vertexSet, e.To)
			}

			source := d.getVertexOfSuperVertex(superVertexToVertex, *SV)

			targetSV := d.getSuperVertexOfVertexSet(SVLookup, &namingSVId, vertexSet)
			target := d.getVertexOfSuperVertex(superVertexToVertex, *targetSV)

			if len(p.edges) == 1 {
				d.CurrGraph.CopyEdge(source, target, *p.edges[0])
			} else {
				//resolve any conflicts
				//TODO create aggregate edges
				print("Partially Implemented merging of edges!\n")

				aggregateEdge := d.CurrGraph.CreateEdge(source, target, p.edges[0].C.Copy())

				//added all interactions

				for _, e := range p.edges {
					aggregateEdge.LoopInteractions.AddInteractions(*e.LoopInteractions)
					aggregateEdge.LoopConditions.AddAll(e.LoopConditions.GetConditions())
				}

				aggregateEdge.LoopInteractions.Solve()
			}

			if !visitedSuperVertices[targetSV.id] {
				visitedSuperVertices[targetSV.id] = true
				SVStack.Push(targetSV)
			}
		}
	}

	//delete all original vertices
	for VID := range originalVertices {
		d.CurrGraph.RemoveVertexFromID(VID)
	}

	return nil
}
