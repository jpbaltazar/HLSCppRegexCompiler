package Builder

type SuperVertexStack struct {
	a []*SuperVertex
}

func NewSuperVertexStack() *SuperVertexStack {
	return &SuperVertexStack{
		make([]*SuperVertex, 0),
	}
}

func (stack *SuperVertexStack) Push(subGraph *SuperVertex) {
	stack.a = append(stack.a, subGraph)
}

func (stack *SuperVertexStack) Pop() *SuperVertex {
	subGraph := stack.a[len(stack.a)-1]
	stack.a = stack.a[:len(stack.a)-1]
	return subGraph
}

func (stack *SuperVertexStack) Peek() *SuperVertex {
	return stack.a[len(stack.a)-1]
}

func (stack *SuperVertexStack) Size() int {
	return len(stack.a)
}
