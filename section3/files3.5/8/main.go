package main

import (
	"bufio"
	"os"
	"strconv"
)

/*
Задача состоит в следующем: на стандартный ввод подаются целые числа в диапазоне 0-100,
каждое число подается на стандартный ввод с новой строки (количество чисел не известно).
Требуется прочитать все эти числа и вывести в стандартный вывод их сумму.
*/

func main() {
	var sum int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		num, errFromParse := strconv.Atoi(s)
		if errFromParse != nil {
			panic(errFromParse)
		}
		sum += num
	}
	os.Stdout.WriteString(strconv.Itoa(sum))
}
