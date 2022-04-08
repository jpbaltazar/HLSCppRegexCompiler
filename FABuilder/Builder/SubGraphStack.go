package Builder

type SubGraphStack struct {
	a []*SubGraph
}

func NewSubGraphStack() *SubGraphStack {
	return &SubGraphStack{
		make([]*SubGraph, 0),
	}
}

func (stack *SubGraphStack) Push(subGraph *SubGraph) {
	stack.a = append(stack.a, subGraph)
}

func (stack *SubGraphStack) Pop() *SubGraph {
	subGraph := stack.a[len(stack.a)-1]

	stack.a = stack.a[:len(stack.a)-2]

	return subGraph
}

func (stack *SubGraphStack) Peek() *SubGraph {
	return stack.a[len(stack.a)-1]
}

func (stack *SubGraphStack) Size() int {
	return len(stack.a)
}
