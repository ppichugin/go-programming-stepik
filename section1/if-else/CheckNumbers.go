package main

import "fmt"

/*
	По данному трехзначному числу определите, все ли его цифры различны
	Формат входных данных
	На вход подается одно натуральное трехзначное число.

	Формат выходных данных
	Выведите "YES", если все цифры числа различны, в противном случае - "NO".
*/

func main() {
	var num int32
	fmt.Scan(&num)

	var secondDigit = num % 100
	var firstDigit = (num - secondDigit) / 100
	var thirdDigit = secondDigit % 10
	secondDigit = secondDigit / 10

	//fmt.Println("=====")
	//fmt.Println(firstDigit)
	//fmt.Println(secondDigit)
	//fmt.Println(thirdDigit)

	if firstDigit != secondDigit && secondDigit != thirdDigit && thirdDigit != firstDigit {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
