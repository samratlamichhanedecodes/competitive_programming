package main

import (
	"fmt"
	"strings"
)

func main281a() {
	var word string
	fmt.Scan(&word)

	if len(word) > 0 {
		capitalizedWord := strings.ToUpper(string(word[0])) + word[1:]
		fmt.Println(capitalizedWord)
	}
}
