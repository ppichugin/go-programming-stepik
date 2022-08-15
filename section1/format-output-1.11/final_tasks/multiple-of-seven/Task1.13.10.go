package main

import "fmt"

/**
Самое большое число, кратное 7
Найдите самое большее число на отрезке от a до b, кратное 7 .

Входные данные
Вводится два целых числа a и b (a≤b).

Выходные данные
Найдите самое большее число на отрезке от a до b
(отрезок включает в себя числа a и b), кратное 7 , или выведите "NO" - если таковых нет.
*/

func main() {
	var a, b int
	var isFound bool
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	for i := b; i >= a; i-- {
		if i%7 == 0 || i == 0 {
			isFound = true
			fmt.Println(i)
			break
		}
	}
	if !isFound {
		fmt.Println("NO")
	}
}
