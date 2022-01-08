package main

type StackString struct {
	stack []string
	upper string
	deep  int
}

func (s *StackString) Push(value string) {
	s.stack = append(s.stack, value)
	s.upper = s.stack[s.deep]
	s.deep++
}

func (s *StackString) Pop() string {
	s.deep--
	value := s.stack[s.deep]
	s.stack = s.stack[:s.deep]
	if s.deep > 0 {
		s.upper = s.stack[s.deep-1]
	} else {
		s.upper = ""
	}
	return value
}

type StackFloat struct {
	stack []float64
	upper float64
	deep  int
}

func (s *StackFloat) Push(value float64) {
	s.stack = append(s.stack, value)
	s.upper = s.stack[s.deep]
	s.deep++
}

func (s *StackFloat) Pop() float64 {
	s.deep--
	value := s.stack[s.deep]
	s.stack = s.stack[:s.deep]
	if s.deep > 0 {
		s.upper = s.stack[s.deep-1]
	} else {
		s.upper = 0.0
	}
	return value
}
