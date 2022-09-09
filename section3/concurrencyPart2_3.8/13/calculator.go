package main

import "fmt"

/*
Необходимо написать функцию calculator следующего вида:
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int

Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int.
 - в случае, если аргумент будет получен из канала firstChan, в выходной (возвращенный) канал вы должны отправить квадрат аргумента.
 - в случае, если аргумент будет получен из канала secondChan, в выходной (возвращенный) канал вы должны отправить результат умножения аргумента на 3.
 - в случае, если аргумент будет получен из канала stopChan, нужно просто завершить работу функции.
Функция calculator должна быть неблокирующей, сразу возвращая управление.
Ваша функция получит всего одно значение в один из каналов - получили значение, обработали его, завершили работу.
*/

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int)
	go func(ch chan int) {
		defer close(ch)
		select {
		case x := <-firstChan:
			ch <- x * x
		case x := <-secondChan:
			ch <- x * 3
		case <-stopChan:
		}
	}(output)
	return output
}

func main() {
	channelToPow := make(chan int)
	channelToMultiply := make(chan int)
	channelToStop := make(chan struct{})
	res := calculator(channelToPow, channelToMultiply, channelToStop)
	channelToPow <- 3
	//channelToMultiply <- 4
	fmt.Println(<-res)
}
