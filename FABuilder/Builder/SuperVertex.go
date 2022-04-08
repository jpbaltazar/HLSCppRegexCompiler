package Builder

import (
	"sort"
)

type SuperVertex struct {
	id       SuperVertexID
	vertices []*Vertex
}

func NewSuperVertex(id SuperVertexID, vertices []*Vertex) *SuperVertex {
	sv := &SuperVertex{
		id,
		vertices,
	}
	sv.SortV()
	return sv
}

func (SV *SuperVertex) IsEqual(vertexSet []*Vertex) bool {
	sort.Slice(vertexSet, func(i, j int) bool { //sort the incoming vertexSet
		return vertexSet[i].Id < vertexSet[j].Id
	})

	if len(SV.vertices) != len(vertexSet) {
		return false
	} else {
		for i := 0; i < len(SV.vertices); i++ {
			v1 := SV.vertices[0]
			v2 := vertexSet[0]

			if v1.Id != v2.Id {
				return false
			}
		}
		return true
	}
}

func (SV *SuperVertex) SortV() {
	sort.Slice(SV.vertices, func(i, j int) bool {
		return SV.vertices[i].Id < SV.vertices[j].Id
	})
}

func (SV *SuperVertex) GetID() SuperVertexID {
	return SV.id
}

func (SV *SuperVertex) GetVertices() []*Vertex {
	return append(make([]*Vertex, 0), SV.vertices...)
}

func NewSuperVertexLookup() map[int][]*SuperVertex {
	return make(map[int][]*SuperVertex)
}

func (d *DFABuilderListener) getSuperVertexOfVertexSet(superVertexLookup map[int][]*SuperVertex, namingID *SuperVertexID, vertices []*Vertex) *SuperVertex {
	superSetList := superVertexLookup[len(vertices)]

	if superSetList == nil {
		superSetList = make([]*SuperVertex, 1)
		superSetList[0] = NewSuperVertex(*namingID, vertices)
		superVertexLookup[len(vertices)] = superSetList
		*namingID++

		return superSetList[0]
	} else {
		for _, sv := range superSetList {
			if sv.IsEqual(vertices) {
				return sv
			}
		}

		newSV := NewSuperVertex(*namingID, vertices)
		*namingID++
		superVertexLookup[len(vertices)] = append(superSetList, newSV)

		return newSV
	}
}

func (d *DFABuilderListener) getVertexOfSuperVertex(superVertexToVertex map[SuperVertexID]*Vertex, vertex SuperVertex) *Vertex {
	if _, ok := superVertexToVertex[vertex.GetID()]; !ok {
		terminal := false

		for _, v := range vertex.GetVertices() {
			if v.Terminal {
				terminal = true
				break
			}
		}

		superVertexToVertex[vertex.GetID()] = d.CurrGraph.CreateVertex(terminal, false)
	}

	return superVertexToVertex[vertex.GetID()]
}
