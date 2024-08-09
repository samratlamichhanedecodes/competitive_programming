package main

import (
	"fmt"
	"strings"
)

func main(){
	var input string

	fmt.Scan(&input)

	input = strings.ToLower(input)


	var newString string
	for _, char := range input {
		charIsVowel := char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'y'
		if !charIsVowel {
			newString += "." + string(char)
		}
	}
	fmt.Println(newString)
}