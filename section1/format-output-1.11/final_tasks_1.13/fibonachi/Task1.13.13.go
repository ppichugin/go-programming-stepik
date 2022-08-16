package main

import "fmt"

/**
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является,
то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.
*/

func fibonacci() func() int {
	a := 1
	b := 1
	return func() int {
		a, b = b, a+b
		return b - a
	}
}

func main() {
	var num int
	fmt.Scanln(&num)
	f := fibonacci()
	for i := 1; ; i++ {
		fib := f()
		if fib == num {
			fmt.Println(i)
			break
		} else if fib > num {
			fmt.Println("-1")
			break
		}
	}
}
