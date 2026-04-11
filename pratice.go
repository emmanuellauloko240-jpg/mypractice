/*package main

import (
	"fmt"
	"strings"
)

func Up(s string) []string {
	words := strings.Fields(s)
	fmt.Println(words)
	fmt.Println(len(words))
	for i, word := range words {
		if word == "(up)" {
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}

func main() {
//	fmt.Println(rune(','))
	fmt.Println(Up("Hello, World! (up)"))
}*/



package main

import (
	"fmt"
	"strings"
)

func functionName(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" {
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}
func main() {
	//word := []int{1, 2, 3, 4, 5}
	fmt.Println(functionName([]string{"john", "paul", "(cap)"}))
}