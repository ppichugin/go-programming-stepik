package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

/**
 * Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
 * Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и буквы латинского алфавита.
 * На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
 */

func main() {
	var text string
	fmt.Scanln(&text)
	rs := []rune(text)
	var res = "Ok"
	length := utf8.RuneCountInString(text)
	if length < 5 {
		res = "Wrong password"
	} else {
		for _, elem := range rs {
			if !(unicode.IsDigit(elem) || unicode.Is(unicode.Latin, elem)) {
				res = "Wrong password"
				break
			}
		}
	}
	fmt.Println(res)
}
