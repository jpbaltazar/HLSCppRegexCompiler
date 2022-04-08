package Builder

type VertexStack struct {
	a []*Vertex
}

func NewVertexStack() *VertexStack {
	return &VertexStack{
		make([]*Vertex, 0),
	}
}

func (stack *VertexStack) Push(vertex *Vertex) {
	stack.a = append(stack.a, vertex)
}

func (stack *VertexStack) Pop() *Vertex {
	subGraph := stack.a[len(stack.a)-1]

	stack.a = stack.a[:len(stack.a)-1]

	return subGraph
}

func (stack *VertexStack) Peek() *Vertex {
	return stack.a[len(stack.a)-1]
}

func (stack *VertexStack) Size() int {
	return len(stack.a)
}
