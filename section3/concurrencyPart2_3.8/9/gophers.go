package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix()) // Настраиваем рандом так, чтобы он был разным для каждого вызова
}

func sleepyGopher(id int, c chan int) {
	duration := time.Duration(rand.Intn(4000)) * time.Millisecond
	fmt.Printf("gopher %d sleep for %v\n", id, duration)
	time.Sleep(duration)
	c <- id
}

func main() {
	timeout := time.After(2 * time.Second)

	c := make(chan int, 5)

	/**
	Горутины для гоферов нужно создать заранее. Если делать это в for вместе с select, то select будет
	блокировать дальнейшее исполнение и создание cледующей горутины
	**/

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	for i := 0; i < 5; i++ {
		select { // Оператор select
		case gopherID := <-c: // Ждет, когда проснется гофер
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case <-timeout: // Ждет окончания времени
			fmt.Println("my patience ran out")

			return // Сдается и возвращается
		}
	}
}
