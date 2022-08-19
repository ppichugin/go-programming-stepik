package main

import "fmt"

/*
 * Поменяйте местами значения переменных на которые ссылаются указатели.
 * После этого переменные нужно вывести.
 */

func main() {
	var a = 2
	var b = 4
	test(&a, &b)
}

func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Printf("%d %d", *x1, *x2)
}
