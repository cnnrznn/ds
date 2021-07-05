package stack

type Stack struct {
	Size int
	Data []interface{}
}

func (s *Stack) Push(data interface{}) {
	s.Size++
	s.Data = append(s.Data, data)
}

func (s *Stack) Pop() interface{} {
	s.Size--
	tmp := s.Data[s.Size]
	s.Data = s.Data[:s.Size]
	return tmp
}
