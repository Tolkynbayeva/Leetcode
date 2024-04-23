package main

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return -1
	}
	index := len(*s) - 1
	elem := (*s)[index]
	*s = (*s)[:index]
	return elem
}

func (s *Stack) Top() int {
	elem := (*s)[len(*s)-1]
	return elem
}

func countStudents(students []int, sandwiches []int) int {
	var stack Stack
	for _, v := range students {
		stack.Push(v)
	}
	count := 0
	for len(stack) > 0 {
		if stack.Top() == sandwiches[0] {
			stack.Pop()
			sandwiches = sandwiches[1:]
			count = 0
		} else {
			if count == len(stack) {
				break
			}
			stopTop := stack.Pop()
			stack = append([]int{stopTop}, stack...)
			count++
		}
	}
	return len(stack)
}
