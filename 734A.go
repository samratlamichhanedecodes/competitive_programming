package main

import (
	"fmt"
)

func main734a() {
	var n uint
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	var antonpoints, danikpoints uint

	for _, char := range s {
		if char == 'A' {
			antonpoints++
		} else {
			danikpoints++
		}
	}

	if antonpoints > danikpoints {
		fmt.Println("Anton")
	} else if danikpoints > antonpoints {
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}
}
