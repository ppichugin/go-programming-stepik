package main

import (
	"fmt"
	"math"
)

func main() {
	var num int32
	fmt.Scan(&num)
	if num < 10 {
		fmt.Println(num)
		return
	}

	for i := 4; i >= 1; i-- {
		res := num / int32(math.Pow(10, float64(i)))
		if res != 0 {
			fmt.Println(res)
			return
		}
	}
}
