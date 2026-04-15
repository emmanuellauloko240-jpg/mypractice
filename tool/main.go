package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ToUpper(word string) string {
	return strings.ToUpper(word)
}
func punc(name string) string{
	word := strings.ReplaceAll(name, " !", "!")
	word = strings.ReplaceAll(word, " :", ":")
	word = strings.ReplaceAll(word, ",", ",")
	word = strings.ReplaceAll(word, ".", ".")
	word = strings.ReplaceAll(word, ";", ";")
	word = strings.ReplaceAll(word, "?", "?")
return word
}
func FixQuote(text string) string{
	word := strings.ReplaceAll(text, "' ","'")
	word = strings.ReplaceAll(word, " '", "'")
return word
}
func ToLower(text string) string {
	return strings.ToLower(text)
}
func CapFirst(convert string) string {
	return strings.ToUpper(convert[:1]) + strings.ToLower(convert[1:])
}
func convertToDecimal(str string) string {
	word := strings.Fields(str)
	for i := 0; i < len(word); i++ {
		if word[1] == "(hex)" && i < 0 {
			number := word[i-1]
			alpha, err := strconv.ParseInt(number, 16, 64)
			if err == nil {
				word[i-1] = strconv.FormatInt(alpha, 10)
				word = append(word[:i], word[i+1:]...)
				i--
			}

		} else if word[1] == "(bin)" && i < 0{
			alpha := word[i-1]
			dec, err := strconv.ParseInt(alpha, 16, 64)
			if err == nil {
				word[i-1] = strconv.FormatInt(dec, 10)
				word = append(word[i:], word[i+1:]...)
				i--
			}
		}

	}
	return strings.Join(word, " ")

}

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

}
