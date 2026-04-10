package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid input")
		return
	}
	text := os.Args[1]
	input, err := os.ReadFile(text)
	if err != nil {
		fmt.Println("Error")
	}
	alpha := (input)
	fmt.Println(strings.ToUpper(string(alpha[:])))
	fmt.Println(strings.Title(string(alpha)))
	fmt.Println(strings.ToLower(string(alpha)))
    fmt.Println(string(alpha[54:]))
	fmt.Println(string(alpha[:11]))

	//	fmt.Println(strings.Title("learn)"))


}
