package main

import (
	"fmt"
)

func main() {
	var num int32
	fmt.Scan(&num)
	var digits [6]int32
	for i := 0; i <= 5; i++ {
		digits[i] = num % 10
		num = num / 10
	}
	var sumLeft = digits[0] + digits[1] + digits[2]
	var sumRight = digits[3] + digits[4] + digits[5]
	if sumLeft == sumRight {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
