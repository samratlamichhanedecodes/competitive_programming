package main

import (
	"fmt"
)

func main266a() {
	var numberOfStones int
	var stones string

	fmt.Scan(&numberOfStones)
	fmt.Scan(&stones)

	var index int
	var removeCount int
	for index < numberOfStones-1 {
		if stones[index] == stones[index+1] {
			removeCount++
			index++
		} else {
			index++
		}
	}
	fmt.Println(removeCount)
}
