package main

import "fmt"

/**
Напишите "функцию голосования", возвращающую то значение (0 или 1),
которое среди значений ее аргументов x, y, z встречается чаще.

Входные данные
Вводится 3 числа - x, y и z (x, y и z равны 0 или 1).
*/

func main() {
	fmt.Println(vote(0, 0, 1))
}

func vote(x int, y int, z int) int {
	if ((x == y || x == z) && x == 0) || y == 0 && z == 0 {
		return 0
	}
	return 1
}
