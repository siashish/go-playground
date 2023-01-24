package main

import "fmt"

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(i int) {
	*s = append(*s, i)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {
	var stack Stack

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y {
			fmt.Println(x)
		}
	}
}
