package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

/**
 * На вход подается строка. Нужно определить, является ли она правильной или нет.
 * Правильная строка начинается с заглавной буквы и заканчивается точкой.
 * Если строка правильная - вывести Right иначе - вывести Wrong
 */

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rs := []rune(text)
	firstLetter := rs[0]
	lastLetter := rs[len(rs)-3] //lastLetter := rs[len(rs)-1] //forStepik
	isUpper := unicode.IsUpper(firstLetter)
	isEndDotted := lastLetter == '.'
	var res string
	if isUpper && isEndDotted {
		res = "Right"
	} else {
		res = "Wrong"
	}
	fmt.Println(res)
}
