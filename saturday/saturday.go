package main

import (
	"fmt"
	"os"
	
)
func main (){
	if len(os.Args) !=3{
		fmt.Println("invalid input")
		return
	}
	
	inputFile := os.Args[1]
	outputFile:= os.Args[2]

	word, err := os.ReadFile(inputFile)
	if err != nil{
		fmt.Println("Error occured")
			return

	}
output := string(word)
	 err = os.WriteFile(outputFile, []byte(output), 0644)	
	 if err != nil{
			fmt.Println("Error while reading file")
	 }

	
}