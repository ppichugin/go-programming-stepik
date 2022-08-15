package main

import (
	"fmt"
	"math"
)

/**
По данному числу N распечатайте все целые значения степени двойки,
не превосходящие N, в порядке возрастания.
*/

func main() {
	var num int
	fmt.Scanln(&num)
	for i := 0; ; i++ {
		res := int(math.Pow(2, float64(i)))
		if res > num {
			break
		}
		fmt.Printf("%d ", res)
	}
}
