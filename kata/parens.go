package kata

import (
	"container/list"
	"fmt"
	"strings"
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
		err := fmt.Errorf("Cannot pop from an empty stack\n")
		return err, ""
	}
	return nil, s.stack.Remove(v).(string)
}

func (s *StringStack) Front() (error, string) {
	v := s.stack.Front()
	if v == nil {
		err := fmt.Errorf("Cannot pop from an empty stack\n")
		return err, ""
	}
	return nil, v.Value.(string)
}

func (s *StringStack) Size() int {
	return s.stack.Len()
}

func ValidParentheses(parens string) bool {
	s := NewStringStack()
	p := strings.Split(parens, "")

	for _, v := range p {
		if v != "(" && v != ")" {
			err := fmt.Errorf("%v is not a valid parenthesis", v)
			fmt.Println(err)
		}

		if v == "(" {
			s.Push(v)
		}
		if v == ")" {
			err, f := s.Front()

			if err != nil {
				fmt.Println(err.Error())
			}

			if len(f) == 0 {
				return false
			}

			if f == "(" {
				err, _ := s.Pop()
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}

	return s.Size() == 0
}
