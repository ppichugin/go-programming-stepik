package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Данная задача в основном ориентирована на изучение типа bufio.Reader, поскольку этот тип позволяет
считывать данные постепенно.

В тестовом файле содержится длинный ряд чисел, разделенных символом ";". Требуется найти, на какой позиции находится
число 0 и указать её в качестве ответа. Требуется вывести именно позицию числа, а не индекс
(то-есть порядковый номер, нумерация с 1).
*/

func main() {
	file, err := os.Open("section3/files3.5/13/task.data")
	if err != nil {
		panic(err)
	}
	rd := bufio.NewReader(file)
	position := 0
	for {
		str, err2 := rd.ReadString(';')
		if err2 != nil {
			panic(err2)
		}
		position++
		if str == "0;" {
			fmt.Println(position)
			break
		}
	}
	file.Close()
}
