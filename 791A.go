package main

import (
	"fmt"
	"math"
)

func main791a() {
	var input1, input2 uint

	fmt.Scan(&input1)
	fmt.Scan(&input2)

	years := 1
	for float64(input1)*math.Pow(3, float64(years)) <= float64(input2)*math.Pow(2, float64(years)) {
		years++
	}

	fmt.Println(years)
}
