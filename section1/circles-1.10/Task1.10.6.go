package main

import "fmt"

func main() {
	var n int // qty of numbers
	var c int // кратное c
	var d int // но НЕ кратное d
	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)
	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}
}
