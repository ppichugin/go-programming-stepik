package main

import (
	"fmt"
	"math"
)

/**
 * На вход подаются a и b - катеты прямоугольного треугольника.
 * Нужно найти длину гипотенузы
 */

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	res := math.Sqrt(float64((a * a) + (b * b)))
	fmt.Println(res)
}
