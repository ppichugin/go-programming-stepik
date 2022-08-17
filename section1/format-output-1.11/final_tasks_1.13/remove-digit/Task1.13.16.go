package main

import "fmt"

/**
Из натурального числа удалить заданную цифру.
*/
func main() {
	var digits [24]int64
	var digitToIgnore, number, quantity int
	fmt.Scanln(&number)
	fmt.Scanln(&digitToIgnore)
	for {
		lastDigit := number % 10
		if lastDigit != digitToIgnore {
			digits[quantity] = int64(lastDigit)
			quantity++
		}
		number = number / 10
		if number == 0 {
			break
		}
	}
	for i := quantity - 1; i >= 0; i-- {
		fmt.Print(digits[i])
	}
}
