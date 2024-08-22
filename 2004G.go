package main

import (
	"fmt"
	"strconv"
)

func getSubstringsArray(s string, n int, k int) []string {
	subStrings := make([]string, n-k+1)
	for i := 0; i <= n-k; i++ {
		subStrings[i] = s[i : i+k]
	}
	return subStrings

}

func findMinNumAndIndexInString(s string, stringLength int) (int, int) {
	minimum := 9
	index := -1

	for i := 0; i < stringLength-2; i++ {
		number, err := strconv.Atoi(s[i : i+1])
		if err != nil {
			return 9, -1
		}
		if number < minimum {
			minimum = number
			index = i
		}
	}

	return minimum, index
}

func findMinlengthPossible(s string, stringLength int) int {
	minLength := 0
	if stringLength&1 == 0 { //length is even

		for i := 0; i < stringLength-1; i += 2 {
			number, err := strconv.Atoi(s[i : i+1]) // Convert substring to integer
			if err != nil {
				return -1
			}
			minLength += number
		}

	} else {
		_, minIndex := findMinNumAndIndexInString(s, stringLength)

		for i := 0; i < stringLength-1; i++ {
			number, err := strconv.Atoi(s[i : i+1])
			if err != nil {
				return 0
			}
			if i == minIndex {
				i++
				minLength += number * 2
			} else {
				minLength += number
			}

		}
	}
	firstNumber, err := strconv.Atoi(s[0:1])
	if err != nil {
		return -1
	}
	if (firstNumber * (stringLength - 1)) < minLength {
		minLength = (firstNumber * (stringLength - 1))
	}

	return minLength
}

func main() {

	var n, k int

	var s string
	fmt.Scan(&n)
	fmt.Scan(&k)
	fmt.Scan(&s)

	subStrings := getSubstringsArray(s, n, k)

	minLengthArray := make([]int, n-k+1)

	for i, value := range subStrings {
		minLengthArray[i] = findMinlengthPossible(value, k)

	}

	for _, length := range minLengthArray {
		fmt.Println(length)
	}

}
