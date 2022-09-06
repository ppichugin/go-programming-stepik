package main

import (
	"fmt"
	"strings"
)

/*
на стандартный ввод вы получаете строку, состоящую ровно из 10 цифр: 0 или 1 (порядок 0/1 случайный).
Ваша задача считать эту строку любым возможным способом и создать на основе этой строки объект
объявленного вами на первом этапе типа: надеюсь, вы понимаете, что строка символизирует
емкость батарейки: 0 - это "опустошенная" часть, а 1 - "заряженная".
*/

func main() {
	var capacity string
	fmt.Scanln(&capacity)
	batteryForTest := &Indicator{charge: capacity}
	fmt.Println(batteryForTest)
}

type Indicator struct {
	charge string
}

func (b *Indicator) String() string {
	num0 := strings.Count(b.charge, "0")
	num1 := strings.Count(b.charge, "1")
	result := "[" + strings.Repeat(" ", num0) + strings.Repeat("X", num1) + "]"
	return result
}
