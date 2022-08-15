package main

import "fmt"

func main() {
	var n, num, qty int
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		fmt.Scanln(&num)
		if num == 0 {
			qty++
		}
	}
	fmt.Println(qty)
}
