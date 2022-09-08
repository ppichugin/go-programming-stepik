package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).
Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между
меньшей и большей датами.

Sample Input:
13.03.2018 14:00:15,12.03.2018 14:00:15

Sample Output:
24h0m0s
*/

func main() {
	var inputTime string
	nr := bufio.NewReader(os.Stdin)
	inputTime, _ = nr.ReadString('\n')
	inputTime = strings.TrimRight(inputTime, "\n")
	inputTime = strings.TrimRight(inputTime, "\r")
	parsedTime := strings.Split(inputTime, ",")
	layout := "02.01.2006 15:04:05"
	var parsedTimeArray [2]time.Time
	var err error
	for i := 0; i <= 1; i++ {
		parsedTimeArray[i], err = time.Parse(layout, parsedTime[i])
		if err != nil {
			fmt.Println(err)
		}
	}
	t1 := parsedTimeArray[0]
	t2 := parsedTimeArray[1]
	var diff time.Duration
	if t1.Unix() > t2.Unix() {
		diff = t1.Sub(t2)
	} else {
		diff = t2.Sub(t1)
	}
	fmt.Println(diff)
}
