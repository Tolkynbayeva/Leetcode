package main

import (
	"strings"
)

func simplifyPath(path string) string {
	stack := []string{}
	newPath := strings.Split(path, "/")
	for _, v := range newPath {
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v != "." && v != "" {
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}
