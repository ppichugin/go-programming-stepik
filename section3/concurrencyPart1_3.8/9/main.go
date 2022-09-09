package main

import (
	"fmt"
)

/*
Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз,
добавив к ней пробел.
Функция должна называться task2().
*/

func main() {
	chanel := make(chan string, 5)
	task2(chanel, "5")
	fmt.Println(<-chanel)
}

//func task2(c chan string, str string) {
//	c <- strings.Repeat(str+" ", 5)
//}

// if package strings not allowed
func task2(c chan string, str string) {
	for i := 0; i < 5; i++ {
		c <- str + " "
	}
}
