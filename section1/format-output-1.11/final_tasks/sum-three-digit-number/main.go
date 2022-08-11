package main

import (
	"fmt"
)

//sum of digits
func main() {
	var num, sum, res int32
	fmt.Scan(&num)
	res = num / 100
	sum = sum + res

	res = num - (res * 100)
	lastDigit := res % 10
	sum += lastDigit

	res = (res - lastDigit) / 10
	sum += res

	fmt.Println(sum)
}
