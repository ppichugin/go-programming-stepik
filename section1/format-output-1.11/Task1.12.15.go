package main

import "fmt"

// print out quantity of positive numbers
func main() {
	array := [100]int{}
	var qtyOfNums uint8
	var qtyOfPositiveNums int
	fmt.Scan(&qtyOfNums)

	for i := 0; i < int(qtyOfNums); i++ {
		fmt.Scan(&array[i])
		if array[i] > 0 {
			qtyOfPositiveNums++
		}
	}
	fmt.Println(qtyOfPositiveNums)
}
