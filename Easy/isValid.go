package main

type Stack []rune

func (s *Stack) Push(val rune) {
	*s = append(*s, val)
}

func (s *Stack) Pop() rune {
	if len(*s) == 0 {
		return 0
	}
	index := len(*s) - 1
	elem := (*s)[index]
	*s = (*s)[:index]
	return elem
}

func isValid(s string) bool {
	var stack Stack
	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack.Push(v)
		} else {
			if len(stack) == 0 || stack.Pop() != m[v] {
				return false
			}
		}
	}
	return len(stack) == 0
}
