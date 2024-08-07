package main

import (
	"fmt"
	"strings"
)

func main(){
	var summationStr string
	fmt.Scan(&summationStr)
	numbersToArrange := strings.Split(summationStr, "+")

	var ones, twos, threes string
	for _, number := range numbersToArrange{
		switch number {
		case "1":
			if len(ones) == 0 {
				ones = number
			}else{
				ones = ones + "+1"
			}
		case "2":
			if len(twos) == 0 {
				twos = number
			}else{
				twos = twos + "+2"
			}
		case "3":
			if len(threes) == 0 {
				threes = number
			}else{
				threes = threes + "+3"
			}
		}

	}
	var newString string

	if len(ones) == 0 {
		if len(twos) == 0 {
			if(len(threes) == 0){
				newString = ""
			}else{
				newString = threes
			}
		}else{
			if(len(threes) == 0){
				newString = twos 
			}else{
				newString = twos + "+" + threes
			}
		}

	}else{
		if len(twos) == 0 {
			if(len(threes) == 0){
				newString = ones
			}else{
				newString = ones + "+" + threes
			}
		}else{
			if(len(threes) == 0){
				newString = ones + "+" + twos 
			}else{
				newString = ones + "+" + twos + "+" + threes
			}
		}

	}

	fmt.Println(newString)
}