package Builder

type EdgeStack struct {
	a []*Edge
}

func NewEdgeStack() *EdgeStack {
	return &EdgeStack{
		make([]*Edge, 0),
	}
}

func (stack *EdgeStack) Push(subGraph *Edge) {
	stack.a = append(stack.a, subGraph)
}

func (stack *EdgeStack) Pop() *Edge {
	subGraph := stack.a[len(stack.a)-1]

	stack.a = stack.a[:len(stack.a)-1]

	return subGraph
}

func (stack *EdgeStack) Peek() *Edge {
	return stack.a[len(stack.a)-1]
}

func (stack *EdgeStack) Size() int {
	return len(stack.a)
}
