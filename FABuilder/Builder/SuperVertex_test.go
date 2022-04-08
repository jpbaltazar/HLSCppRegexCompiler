package Builder

import "testing"

func TestSort(t *testing.T) {
	sV := SuperVertex{
		0,
		[]*Vertex{
			&Vertex{
				1,
				nil,
				nil,
				false,
				false,
			},
			&Vertex{
				0,
				nil,
				nil,
				false,
				false,
			},
		},
	}

	sV.SortV()

	for _, v := range sV.vertices {
		print(v.Id)
		print("\n")
	}
}
