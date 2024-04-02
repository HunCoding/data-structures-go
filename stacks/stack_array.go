package main

type Stack struct {
	items []string
}

func (s *Stack) push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() string {
	if len(s.items) == 0 {
		return ""
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) peek() string {
	if len(s.items) == 0 {
		return ""
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}
