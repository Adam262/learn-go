package kata

import (
	"container/list"
	"fmt"
)

type StringStack struct {
	stack *list.List
}

func NewStringStack() *StringStack {
	l := list.New()
	return &StringStack{l}
}

func (s *StringStack) Push(str string) {
	s.stack.PushFront(str)
}

func (s *StringStack) Pop() (error, string) {
	v := s.stack.Front()
	if v == nil {
		err := fmt.Errorf("Cannot pop from an empty stack")
		fmt.Print(err)
		return err, ""
	}
	return nil, s.stack.Remove(v).(string)
}

func (s *StringStack) Size() int {
	return s.stack.Len()
}
