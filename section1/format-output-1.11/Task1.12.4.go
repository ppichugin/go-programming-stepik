package main

import "fmt"

func main() {
	var workArray [10]uint8
	var i, j, a, c uint8

	for i = 0; i < 10; i++ {
		fmt.Scan(&a)
		workArray[i] = a
	}

	for c = 0; c < 3; c++ {
		fmt.Scan(&i, &j)
		workArray[i], workArray[j] = workArray[j], workArray[i]
	}

	for _, element := range workArray {
		fmt.Print(element, " ")
	}
	fmt.Print("typed ok")
}
