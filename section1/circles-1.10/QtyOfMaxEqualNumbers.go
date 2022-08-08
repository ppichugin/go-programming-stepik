// 1.10 Циклы. 4 из 10
package main

import "fmt"

func main() {
	var a int      // input number
	var maxNum = 0 // max number in sequence
	var qty = 0    // qty of identical max numbers

	for true {
		fmt.Scan(&a)
		if a == 0 {
			break
		}
		if a == maxNum {
			qty = qty + 1
			continue
		}
		if a > maxNum {
			maxNum = a
			qty = 1
		}
	}
	fmt.Println(qty)

}
