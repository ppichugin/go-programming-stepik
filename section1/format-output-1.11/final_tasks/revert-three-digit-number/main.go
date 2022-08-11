package main

import "fmt"

// Дано трехзначное число. Переверните его, а затем выведите.
func main() {
	var num, reverseNumber, res int32
	fmt.Scan(&num)
	first := num / 100
	res = num - (first * 100)
	third := res % 10
	second := (res - third) / 10
	reverseNumber = third*100 + second*10 + first
	fmt.Println(reverseNumber)
}
