package main

import (
	"fmt"
)

func isCharacterInTheString(stringToFindCharIn string, character rune) bool {

	for _, char := range stringToFindCharIn{
		if(character == char){
			return true
		}
	}
	return false

}

func main(){
	var userName string
	fmt.Scan(&userName)

	var userNameWithUniqueChars string
	for _, char := range userName{
		if !isCharacterInTheString(userNameWithUniqueChars, char) {
			userNameWithUniqueChars = userNameWithUniqueChars + string(char)
		}
	}

	if (len(userNameWithUniqueChars) & 1) == 1 {
		fmt.Println("IGNORE HIM!")
	}else{
		fmt.Println("CHAT WITH HER!")
	}
}