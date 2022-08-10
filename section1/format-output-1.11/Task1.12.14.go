package main

import "fmt"

// print out all odd array's index
func main() {
	array := [100]int{}
	var qtyOfNums uint8
	fmt.Scan(&qtyOfNums)

	for i := 0; i < int(qtyOfNums); i++ {
		fmt.Scan(&array[i])
	}

	for i := 0; i < int(qtyOfNums); i += 2 {
		fmt.Printf("%d ", array[i])
	}
}
