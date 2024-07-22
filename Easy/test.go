package main

import "fmt"

func print(n int) {
	if n > 0 {
		fmt.Println(n)
		print(n-1)
	}
}

func main() {
	n := 5
	print(n)
}
