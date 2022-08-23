package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/**
 * Дается строка. Нужно удалить все символы, которые встречаются более одного раза
 * и вывести получившуюся строку.
 */

func main() {
	var text string
	fmt.Scanln(&text)
	rs := []rune(text)
	var res []string
	for i := 0; i < utf8.RuneCountInString(text); i++ {
		elem := string(rs[i])
		if strings.Count(text, elem) <= 1 {
			res = append(res, elem)
		}
	}
	fmt.Println(strings.Join(res, ""))
}
