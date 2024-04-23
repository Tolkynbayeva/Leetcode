package main

import (
	"fmt"
	"math"
	"strconv"
)

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Top() int {
	index := len(*s) - 1
	elem := (*s)[index]
	return elem
}

func (s *Stack) PrevTop() int {
	index := len(*s) - 2
	elem := (*s)[index]
	return elem
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return 0
	}
	index := len(*s) - 1
	elem := (*s)[index]
	*s = (*s)[:index]
	return elem
}

func evalRPN(tokens []string) int {
	var stack Stack

	for _, v := range tokens {
		switch v {
		case "+":
			last := stack.Pop()
			lastPrev := stack.Pop()
			stack.Push(lastPrev + last)
		case "-":
			last := stack.Pop()
			lastPrev := stack.Pop()
			stack.Push(lastPrev - last)
		case "*":
			last := stack.Pop()
			lastPrev := stack.Pop()
			stack.Push(lastPrev * last)
		case "/":
			last := stack.Pop()
			lastPrev := stack.Pop()
			result := float64(lastPrev) / float64(last)
			stack.Push(int(math.Trunc(result)))
		default:
			num, _ := strconv.Atoi(v)
			stack.Push(num)
		}
	}
	return stack.Top()
}
func main() {
	fmt.Println(evalRPN([]string{"4","13","5","/","+"})) //6
	fmt.Println(evalRPN([]string{"-78", "-33", "196", "+", "-19", "-", "115", "+", "-", "-99", "/", "-18", "8", "*", "-86", "-", "-", "16", "/", "26", "-14", "-", "-", "47", "-", "101", "-", "163", "*", "143", "-", "0", "-", "171", "+", "120", "*", "-60", "+", "156", "/", "173", "/", "-24", "11", "+", "21", "/", "*", "44", "*", "180", "70", "-40", "-", "*", "86", "132", "-84", "+", "*", "-", "38", "/", "/", "21", "28", "/", "+", "83", "/", "-31", "156", "-", "+", "28", "/", "95", "-", "120", "+", "8", "*", "90", "-", "-94", "*", "-73", "/", "-62", "/", "93", "*", "196", "-", "-59", "+", "187", "-", "143", "/", "-79", "-89", "+", "-"})) //165
}
