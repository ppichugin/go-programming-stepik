// задача 1.10 (3)
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a) // кол-во чисел
	var sum int32
	for i := 1; i <= a; i++ {
		fmt.Scan(&b)
		if (b > 9 && b < 100) && b%8 == 0 {
			sum += int32(b)
		}
	}
	fmt.Println(sum)
}
