package main

import (
	"fmt"
)

func main() {
	var berland, birland string

	fmt.Scan(&berland)
	fmt.Scan(&birland)

	birlandLength := len(birland)

	for i, character := range berland {
		if character == rune(birland[birlandLength-i-1]) {
			if i == birlandLength-1 {
				fmt.Println("YES")
				break
			}
			continue
		} else {
			fmt.Println("NO")
			break
		}
	}

}
