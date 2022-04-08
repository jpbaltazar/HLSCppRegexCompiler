package FABuilder

type CaptureGroupStack struct {
	a []string
}

func NewCaptureGroupStack() *CaptureGroupStack {
	return &CaptureGroupStack{
		make([]string, 0),
	}
}

func (stack *CaptureGroupStack) Push(str string) {
	stack.a = append(stack.a, str)
}

func (stack *CaptureGroupStack) Pop() string {
	subGraph := stack.a[len(stack.a)-1]

	stack.a = stack.a[:len(stack.a)-1]

	return subGraph
}

func (stack *CaptureGroupStack) GetAll() []string {
	return stack.a
}
