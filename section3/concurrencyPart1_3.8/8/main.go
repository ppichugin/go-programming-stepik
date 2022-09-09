package main

import "fmt"

/*
Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал.
Функция должна называться task().
*/

func main() {
	chanel := make(chan int, 5)
	task(chanel, 5)
	fmt.Println(<-chanel)
}

func task(c chan int, n int) {
	c <- n + 1
}
