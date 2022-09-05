package main

import (
	"fmt"
	"strconv"
)

/*
Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют
нечетные цифры и цифра 0. Если же получившееся число равно 0, то возвращается число 100.
Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.
*/

func main() {

	fn := func(num uint) uint {
		var res string
		stringFromNum := strconv.Itoa(int(num))
		for _, x := range stringFromNum {
			if x&1 == 0 && x != 48 {
				res += string(x)
			}
		}
		if res == "" {
			return 100
		}
		finalRes, err := strconv.Atoi(res)
		if err != nil {
			panic("ERROR on parse to int")
		}
		return uint(finalRes)
	}

	fmt.Println(fn(71))

}
