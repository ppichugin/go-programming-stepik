package main

import "fmt"

/* ЗАДАНИЕ:
 * Напишите функцию sumInt, получающую переменное число аргументов типа int,
 * и возвращающую количество переданных аргументов и их сумму.
 */

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}

func sumInt(a ...int) (int, int) {
	var length = len(a)
	var sum int
	for _, elem := range a {
		sum += elem
	}
	return length, sum
}
