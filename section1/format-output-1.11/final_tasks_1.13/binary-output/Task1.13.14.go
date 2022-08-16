package main

import "fmt"

/**
Дано натуральное число N. Выведите его представление в двоичном виде.
*/

func main() {
	var num int
	fmt.Scanln(&num)
	fmt.Printf("%b", num)
}
