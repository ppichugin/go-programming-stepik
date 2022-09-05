package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
На стандартный ввод вы получаете 2 таких вещественных числа, в качестве результата требуется вывести частное
от деления первого числа на второе с точностью до четырех знаков после "запятой"
(на самом деле после точки, результат не требуется приводить к исходному формату).

Sample Input:
1 149,6088607594936;1 179,0666666666666

Sample Output:
0.9750

*/

func main() {
	var array [2]float64
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, "\r")
	split := strings.Split(str, ";")
	for i := 0; i < 2; i++ {
		section := split[i]
		section = strings.ReplaceAll(section, " ", "")
		section = strings.ReplaceAll(section, ",", ".")
		parseFloat, err := strconv.ParseFloat(section, 64)
		if err != nil {
			println(err)
			panic("ERROR in ParseFloat: ")
		}
		array[i] = parseFloat
	}
	fmt.Printf("%.4f", array[0]/array[1])
}
