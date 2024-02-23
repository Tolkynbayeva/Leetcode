package main

import (
	"fmt"
	"strings"
)

func interpret(command string) string {
	first := strings.ReplaceAll(command, "()", "o")
	result := strings.ReplaceAll(first, "(al)", "al")
    return result
}


func main() {
	//result := interpret("G()(al)") 
	//result := interpret("G()()()()(al)")
	result := interpret("(al)G(al)()()G")
	fmt.Println(result)
}

