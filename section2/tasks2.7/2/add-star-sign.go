package main

import (
	"fmt"
	"unicode/utf8"
)

/**
 * Дана строка, содержащая только английские буквы (большие и маленькие).
 * Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и
 * после последней символ ‘*’ добавлять не нужно).
 */

func main() {
	var text string
	fmt.Scan(&text)
	for idx, elem := range text {
		switch idx {
		case 0:
			fmt.Print(string(elem))
		case utf8.RuneCountInString(text):
			fmt.Printf("%s*", string(elem))
		default:
			fmt.Printf("*%s", string(elem))
		}
	}
}
