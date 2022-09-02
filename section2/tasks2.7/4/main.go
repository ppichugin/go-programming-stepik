package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

/**
 * На вход подается целое число. Необходимо возвести в квадрат каждую цифру
 * числа и вывести получившееся число.
 * Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81.
 * Дальше 1. Единица в квадрате - 1. В итоге получаем 811181
 */

func main() {
	var digits string
	_, err := fmt.Scan(&digits)
	if err != nil {
		return
	}
	length := utf8.RuneCountInString(digits)
	for i := 0; i < length; i++ {
		num, _ := strconv.Atoi(string(digits[i]))
		fmt.Print(math.Pow(float64(num), 2))
	}
}
