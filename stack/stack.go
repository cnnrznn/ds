package stack

type Stack[T any] struct {
	Size int
	Data []T
}

func (s *Stack[T]) Push(data T) {
	s.Size++
	s.Data = append(s.Data, data)
}

func (s *Stack[T]) Pop() T {
	s.Size--
	tmp := s.Data[s.Size]
	s.Data = s.Data[:s.Size]
	return tmp
}
