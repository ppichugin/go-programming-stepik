package main

import "fmt"

func main() {
	var (
		x, p, y int
		result  = 0
	)
	fmt.Scan(&x) // сумма вклада
	fmt.Scan(&p) // ставка %
	fmt.Scan(&y) // ожидаемая сумма вклада
	for year := 1; ; year++ {
		x = x + (x * p / 100)
		if (x) >= y {
			result = year
			break
		}
	}
	fmt.Println(result)
}
