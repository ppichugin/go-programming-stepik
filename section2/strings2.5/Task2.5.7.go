package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * Если строка является палиндромом - вывести Палиндром иначе - вывести Нет.
 * Все входные строчки в нижнем регистре.
 */

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rs := []rune(text)
	lastIndex := len(rs) - 3 //lastLetter := rs[len(rs)-1] //forStepik
	var res = "Палиндром"
	for i := lastIndex; i >= 0; i-- {
		elem := string(rs[i])
		if elem != string(rs[lastIndex-i]) {
			res = "Нет"
			break
		}
	}
	fmt.Println(res)
}
