package main

import "fmt"

// find max number in array
func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	max := array[0]
	for i, value := range array {
		if array[i] > max {
			max = value
		}
	}
	fmt.Println(max)
}
