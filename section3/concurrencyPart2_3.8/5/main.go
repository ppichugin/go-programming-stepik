package main

import (
	"fmt"
	"sync"
)

/*
необходимо в отдельных горутинах вызвать функцию work() 10 раз и дождаться результатов выполнения вызванных функций.
*/
func main() {
	var wg = new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}
	wg.Wait()
	fmt.Println("All goroutines finished")
}

func work() {
	fmt.Println("working from goroutine")
}
