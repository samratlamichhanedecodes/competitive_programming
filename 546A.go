package main

import (
	"fmt"
)

func main(){
	var k, w int
	var n int64

	fmt.Scan(&k, &n, &w)
	//money for banana if it costs the i.k for the i'th banana you want. cost if you want w bananas
	moneyForBanana := k * (w * (w + 1) / 2)

	if int64(moneyForBanana) > n {
		fmt.Println(int64(moneyForBanana) - n)
	}else{
		fmt.Println(0)
	}

}