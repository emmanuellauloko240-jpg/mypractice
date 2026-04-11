/*package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	text, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error occured while reading file", err)
		return
	}


	word := string(text)
	err = os.WriteFile(outputFile, []byte(word), 0644)
	if err != nil {
		fmt.Println("Error writing content to outputFile", err)
	}

}*/
package main
	