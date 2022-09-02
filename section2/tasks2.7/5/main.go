package main

import (
	"fmt"
	"math"
)

var k, p, v float64

func main() {
	fmt.Scan(&k, &p, &v)
	fmt.Println(T())
}

func T() float64 {
	t1 := W()
	return 6 / t1
}

func W() float64 {
	m1 := M()
	return math.Sqrt(k / m1)
}

func M() float64 {
	return p * v
}
