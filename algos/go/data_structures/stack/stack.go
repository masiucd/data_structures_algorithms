package stack

import "fmt"

type Stackable interface {
	Push(int)           //  which pushes an element onto the stack
	Pop() (int, error)  // which pops off and returns the topmost element of the stack.
	Peek() (int, error) // returns the first value in the stack
}

type Stack struct {
	stack []int
}

func NewStack() *Stack {
	return &Stack{
		stack: []int{},
	}
}

func (s *Stack) Push(val int) {
	s.stack = append(s.stack, val)
}

func (s *Stack) Pop() (int, error) {
	if len(s.stack) == 0 {
		return -1, fmt.Errorf("stack is empty")
	}
	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return val, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.stack) == 0 {
		return -1, fmt.Errorf("stack is empty")
	}
	return s.stack[len(s.stack)-1], nil
}
