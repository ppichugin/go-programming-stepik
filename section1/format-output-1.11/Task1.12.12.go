package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	slice := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scan(&slice[i])
	}
	fmt.Println(slice[3])
}
