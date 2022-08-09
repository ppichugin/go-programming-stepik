package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		num    float64
		result float64
	)
	fmt.Scan(&num)
	result = math.Pow(num, 2) //fmt.Println(result)
	if num <= 0 {
		fmt.Printf("число %2.2f не подходит", num)
	} else if num > 10000 {
		fmt.Printf("%e", num)
		return
	} else {
		fmt.Printf("%.4f", result)
	}
}
