package main

import (
	"fmt"
	"strings"
)

/**
 * Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X.
 * Если подстроки S нет в строке X - вывести -1
 */

func main() {
	var text string
	var subStr string
	fmt.Scanln(&text)
	fmt.Scanln(&subStr)
	index := strings.Index(text, subStr)
	fmt.Println(index)
}
