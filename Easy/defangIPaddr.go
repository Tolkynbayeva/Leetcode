package main

import (
	"strings"
)

func defangIPaddr(address string) string {
	address = strings.ReplaceAll(address, ".", "[.]")
	return address
}

// func main() {
// 	fmt.Println(defangIPaddr("1.1.1.1"))
// }
