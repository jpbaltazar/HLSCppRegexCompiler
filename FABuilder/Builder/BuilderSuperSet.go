package Builder

import (
	"fmt"
	"thesisGoRemake/FABuilder/NumberedLoops"
)

type SuperVertexID uint64

//creates both transitionMap and epsilon set, also return visitedVertices (for later deletion)
func (d *DFABuilderListener) populateStartingSet(start *Vertex) (map[VertexID][]*Edge, map[VertexID][]*Vertex, map[VertexID]bool) {
	transitionMap := make(map[VertexID][]*Edge)        //Map<Vertex, ArrayList<Edge>>
	epsilonSetOfVertex := make(map[VertexID][]*Vertex) //Map<Vertex, Set<Vertex>>

	vertexStack := NewVertexStack()
	visitedVertices := make(map[VertexID]bool)
	vertexStack.Push(start)
	visitedVertices[start.Id] = true

	for vertexStack.Size() > 0 {
		v := vertexStack.Pop()
		print("Popping V")
		print(v.Id)
		print("\n With ")
		print(len(v.Outgoing))
		print(" outgoing edges\n")
		transitionMap[v.Id] = make([]*Edge, 0)

		epsilonSetOfVertex[v.Id] = d.epsilonSet(v)
		for _, e := range v.Outgoing {
			if _, ok := visitedVertices[e.To.Id]; !ok {
				visitedVertices[e.To.Id] = true

				vertexStack.Push(d.CurrGraph.GetVertex(e.To.Id))
			}

			if !e.IsEpsilon() {
				transitionMap[v.Id] = append(transitionMap[v.Id], e)
			}
		}
	}

	return transitionMap, epsilonSetOfVertex, visitedVertices
}

//calls the recursive version
func (d *DFABuilderListener) epsilonSet(vertex *Vertex) []*Vertex {
	m := d.recursiveEpsilonSet(vertex)

	ret := make([]*Vertex, len(m))

	i := 0
	for _, v := range m {
		ret[i] = v
		i++
	}

	return ret
}

//recursively composes the epsilon set
func (d *DFABuilderListener) recursiveEpsilonSet(vertex *Vertex) map[VertexID]*Vertex {
	set := make(map[VertexID]*Vertex)
	set[vertex.Id] = vertex

	for _, e := range vertex.Outgoing {
		if e.IsEpsilon() {
			newSet := d.recursiveEpsilonSet(d.CurrGraph.GetVertex(e.To.Id))
			for id := range newSet {
				set[id] = newSet[id]
			}
		}
	}

	return set
}

func (d *DFABuilderListener) SuperSetSolver(start *Vertex) (newStart, newStop *Vertex) {
	//SuperSet solver
	//take epsilon connected groups of vertices and translate them to a single vertex
	//if an edge of charset 0 and LoopConditions !0, try to connect them to a charset !0 edge

	transitionMap, epsilonSet, visitedVertices := d.populateStartingSet(start)
	//transitionMap -> all non epsilon edges of the vertices
	//epsilonSet -> epsilon group of a vertex
	//visitedVertices -> vertices visited (important for deletion later)

	//we're going to need a lookup of the vertex sets
	//criteria:
	// -number of items
	//		-easy discard
	// -set values
	//		-harder, but if they are sorted by Ids, it can be fairly fast

	//	 Array vs Map?
	//	Ordering for lookup of the values?
	// Array may be better

	//map[uint64] []*SuperVertex
	// size of superVertex -> possible SuperVertices
	// try all the SuperVertices to see if it is the same
	// could work

	superVertexLookUpSize := make(map[int][]*SuperVertex)

	currSVId := SuperVertexID(0)
	superVertexToVertex := make(map[SuperVertexID]*Vertex)

	visitedSuperVertices := make(map[SuperVertexID]bool)

	startSV := d.getSuperVertexOfVertexSet(superVertexLookUpSize, &currSVId, epsilonSet[start.Id])
	sVStack := NewSuperVertexStack()

	newStart = d.getVertexOfSuperVertex(superVertexToVertex, *startSV)

	sVStack.Push(startSV)

	for sVStack.Size() > 0 {
		sV := sVStack.Pop()

		//for all vertices of the super vertex,

		//check for an edge to connect

		for _, v := range sV.GetVertices() {
			for _, e := range transitionMap[v.Id] {
				if !e.C.IsEmpty() {
					sourceEpsilonSet := epsilonSet[e.To.Id]

					conditionedEdgeFound := false
					for _, v2 := range sourceEpsilonSet {
						for _, e2 := range transitionMap[v2.Id] {
							if e2.C.IsEmpty() && !e2.LoopConditions.IsEmpty() {
								conditionedEdgeFound = true

								targetVertexSet := sourceEpsilonSet
								targetVertexSet = append(targetVertexSet, epsilonSet[e.To.Id]...)

								targetSV := d.getSuperVertexOfVertexSet(superVertexLookUpSize, &currSVId, targetVertexSet)

								//c -> corresponding
								//we are building a copy of the subgraph and turning them into a non-Epsilon subgraph
								cSourceVertex := d.getVertexOfSuperVertex(superVertexToVertex, *sV)
								cTargetVertex := d.getVertexOfSuperVertex(superVertexToVertex, *targetSV)

								if cSourceVertex.Terminal {
									//if the conditions are incomplete, n
									if e2.LoopConditions.GetConditions()[0].Type == NumberedLoops.Complete {
										superSetEdge := d.CurrGraph.CreateEpsilonEdge(cSourceVertex, cTargetVertex)

										if e2.LoopConditions.GetConditions()[0].Var.Type == NumberedLoops.Bound {
											superSetEdge.LoopConditions.AddAll(e2.LoopConditions.GetConditions())
										}
									}

									//no need to do so for impossible transitions
									//i.e. i < 0 (incomplete) when i<=1 (complete) has been travelled
								} else {
									superSetEdge := d.CurrGraph.CreateEpsilonEdge(cSourceVertex, cTargetVertex)
									superSetEdge.LoopConditions.AddAll(e2.LoopConditions.GetConditions())
								}

								if _, ok := visitedSuperVertices[targetSV.GetID()]; !ok {
									sVStack.Push(targetSV)
									visitedSuperVertices[targetSV.GetID()] = true
								}
							}
						}
					}

					if !conditionedEdgeFound {
						//no transformation needed, simply add directly
						cSourceVertex := d.getVertexOfSuperVertex(superVertexToVertex, *sV)
						targetSV := d.getSuperVertexOfVertexSet(superVertexLookUpSize, &currSVId, sourceEpsilonSet)

						cTargetVertex := d.getVertexOfSuperVertex(superVertexToVertex, *targetSV)

						if _, ok := cSourceVertex.Outgoing[cTargetVertex.Id]; !ok {
							d.CurrGraph.CopyEdge(cSourceVertex, cTargetVertex, *e)
						} else {
							//if they only differ in how they interact (set/add) with variables, just add those interactions
							edge := cSourceVertex.Outgoing[cTargetVertex.Id]
							if edge.C.EqualTo(*e.C) && edge.LoopConditions.IsEqual(e.LoopConditions) {
								edge.LoopInteractions.AddInteractions(*e.LoopInteractions)
							} else {
								print(fmt.Errorf("Build error! Edge already existed between %d and %d, and it is different from the proposed one...\n", cSourceVertex.Id, cTargetVertex.Id))
							}
						}

						if _, ok := visitedSuperVertices[targetSV.GetID()]; !ok {
							sVStack.Push(targetSV)
							visitedSuperVertices[targetSV.GetID()] = true
						}
					}
				}

			}
		}
	}

	for vID := range visitedVertices {
		d.CurrGraph.RemoveVertexFromID(vID)
	}

	//For clarity this is a Vertex representation of the SuperVertex
	terminalSuperVertices := make(map[VertexID]bool)
	var lastTerminal *Vertex
	for v := range visitedSuperVertices {
		cSuperVertex := superVertexToVertex[v]
		if cSuperVertex.Terminal {
			terminalSuperVertices[cSuperVertex.Id] = true
			lastTerminal = cSuperVertex
		}
	}

	if len(terminalSuperVertices) > 1 {
		totalTerminal := d.CurrGraph.CreateTerminalVertex()

		for terminalID := range terminalSuperVertices {
			d.CurrGraph.CreateEpsilonEdge(d.CurrGraph.GetVertex(terminalID), totalTerminal)
		}

		newStop = totalTerminal
	} else {
		newStop = lastTerminal
	}

	return newStart, newStop
}
