package main

import "fmt"

/**
Написать функцию, которая возвращает мин. из 4-х введенных чисел
*/
func main() {
	fmt.Println(minimumFromFour())
}

func minimumFromFour() int {
	var min, num int
	for i := 0; i < 4; i++ {
		fmt.Scan(&num)
		if i == 0 {
			min = num
			continue
		}
		if num < min {
			min = num
		}
	}
	return min
}
