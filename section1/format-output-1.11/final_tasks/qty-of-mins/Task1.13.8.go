package main

import "fmt"

func main() {
	var n, num, mins, qty int
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		fmt.Scanln(&num)
		if i == 1 {
			mins = num
			qty++
			continue
		}
		if num < mins {
			mins = num
			qty = 1
		} else if num == mins {
			qty++
		}
	}
	fmt.Println(qty)
}
