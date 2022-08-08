package main

import "fmt"

func main() {
	var a, b int32
	fmt.Scan(&a)
	fmt.Scan(&b)
	var sum int32
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Println(sum)
}
