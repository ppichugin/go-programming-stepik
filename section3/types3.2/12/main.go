package main

import (
	"fmt"
	"strconv"
	"unicode"
)

/*
Представьте что вы работаете в большой компании где используется модульная архитектура.
Ваш коллега написал модуль с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные.
Вы же пишете функцию которая считывает две переменных типа string,
а возвращает число типа int64 которое нужно получить сложением этих строк.

Но не было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того, что гоферам платят больше.
Поэтому он решил подшутить над вами и подсунул вам свинью.
Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.

Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа.
Под мусором имеются ввиду лишние символы и спец знаки.
Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes.
*/

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	fmt.Println(adding(a, b))
}

func adding(str1 string, str2 string) int64 {
	var num1, num2 string

	runes := []rune(str1)
	for i := range runes {
		if unicode.IsDigit(runes[i]) {
			num1 += string(runes[i])
		}
	}

	runes = []rune(str2)
	for i := range runes {
		if unicode.IsDigit(runes[i]) {
			num2 += string(runes[i])
		}
	}

	atoi1, err := strconv.Atoi(string(num1))
	if err != nil {
		panic("ERROR on num1")
	}

	atoi2, err := strconv.Atoi(string(num2))
	if err != nil {
		panic("ERROR on num2")
	}

	return int64(atoi1 + atoi2)
}
