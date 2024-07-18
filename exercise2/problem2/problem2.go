package problem2

import (
	"fmt"
)

type Stack struct {
	items []any
}

func (s *Stack) Push(e any) {
	if s.IsEmpty() {
		s.items = make([]any, 0)
	}

	s.items = append(s.items, e)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Pop() (any, error) {
	res, err := s.Peek()

	if err != nil {
		return nil, err
	}

	s.items = s.items[:s.Size()-1]

	return res, nil
}

func (s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}

	res := s.items[s.Size()-1]

	return res, nil
}

func (s *Stack) Size() int {
	return len(s.items)
}
