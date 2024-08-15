package main

import (
	"fmt"
)

func main69A() {
	var rows int
	fmt.Scan(&rows)

	forces := make([][]int, rows)
	for i := range forces {
		forces[i] = make([]int, 3)
	}

	for i := range forces {
		for j := range forces[i] {
			fmt.Scan(&forces[i][j])
		}
	}

	resultant := make([]int, 3)

	for i := range forces {
		for j := range forces[i] {
			resultant[j] += forces[i][j]
		}
	}

	if resultant[0] == 0 && resultant[1] == 0 && resultant[2] == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println(("NO"))
	}

}
