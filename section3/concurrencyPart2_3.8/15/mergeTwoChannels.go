package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Необходимо написать функцию func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).

Описание ее работы:

n раз сделать следующее:
 - прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
 - вычислить f(x1) + f(x2)
 - записать полученное значение в out

Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.
Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.

Формат ввода:
 - количество итераций передается через аргумент n.
 - целые числа подаются через аргументы-каналы in1 и in2.
 - функция для обработки чисел перед сложением передается через аргумент fn.

Формат вывода:
 - канал для вывода результатов передается через аргумент out.
*/

const N = 20

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func(newFn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int) {
		done := make(chan bool, n*2)
		results := make([]*int, n*2)

		go func() {
			nDone := 0
			for <-done {
				for i := nDone; i < n; i++ {
					if results[i] != nil && results[i+n] != nil {
						out <- *results[i] + *results[i+n]
						if nDone++; nDone == n {
							return
						}
					} else {
						break
					}
				}
			}
		}()

		input := func(ch <-chan int, results []*int) {
			for i := 0; i < n; i++ {
				x := <-ch
				go func(i int, x int) {
					result := newFn(x)
					results[i] = &result
					done <- true
				}(i, x)
			}
		}
		go input(in1, results[:n])
		go input(in2, results[n:])

	}(fn, in1, in2, out)
}

func main() {
	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}
