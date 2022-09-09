package main

import "fmt"

/*
Внутри функции main, необходимо в отдельной горутине вызвать функцию work() и дождаться результатов ее выполнения.
Функция work() ничего не принимает и не возвращает.
*/

func main() {
	done := make(chan struct{})
	go func() {
		work()
		close(done)
	}()
	<-done
}

func work() {
	fmt.Println("working from goroutine")
}
