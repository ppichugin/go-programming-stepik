package main

import (
	"fmt"
)

/**
Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1.
Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
Напишите функцию, которая по указанному натуральному n возвращает φn.
*/

func main() {
	fmt.Println(fibonacci(4))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
