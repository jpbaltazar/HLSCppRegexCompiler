package Builder

type Stack[T interface{}] struct {
	a []T
}

func (s *Stack[T]) Push(item T) {
	s.a = append(s.a, item)
}

func (s *Stack[T]) Pop() (item T) {
	item = s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]

	return item
}

func (s *Stack[T]) Peek() T {
	return s.a[len(s.a)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.a)
}
