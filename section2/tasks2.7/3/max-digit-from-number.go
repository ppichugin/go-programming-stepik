package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

/**
 * Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.
 */

func main() {
	var digits string
	_, err := fmt.Scan(&digits)
	if err != nil {
		return
	}
	var res int
	length := utf8.RuneCountInString(digits)
	res, _ = strconv.Atoi(string(digits[0]))
	for i := 1; i < length; i++ {
		num, _ := strconv.Atoi(string(digits[i]))
		if num > res {
			res = num
		}
	}
	fmt.Println(res)
}
