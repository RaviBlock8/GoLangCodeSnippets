package main

import "fmt"

type stack struct {
	stack []int
	top   int
}

func (s *stack) push(value int) bool {
	s.stack = append(s.stack, value)
	s.top += 1
	return true
}

func (s *stack) pop(value int) bool {
	if s.top == 0 {
		fmt.Errorf("Underflow")
		return false
	}
	s.stack = s.stack[0 : len(s.stack)-2]
	s.top -= 1
	return true
}
