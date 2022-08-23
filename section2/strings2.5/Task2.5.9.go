package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/**
 * На вход дается строка, из нее нужно сделать другую строку,
 * оставив только нечетные символы (считая с нуля)
 */

func main() {
	var text string
	fmt.Scanln(&text)
	rs := []rune(text)
	var res []string
	for i := 1; i < utf8.RuneCountInString(text); i += 2 {
		elem := string(rs[i])
		res = append(res, elem)
	}
	fmt.Println(strings.Join(res, ""))
}
