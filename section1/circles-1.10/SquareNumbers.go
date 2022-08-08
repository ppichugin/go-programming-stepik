package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i <= 10; i++ {
		var a = math.Pow(float64(i), 2)
		fmt.Println(a)
	}
}
