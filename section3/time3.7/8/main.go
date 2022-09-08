package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере).
Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64
(наносекунды для целей преобразования равны 0).

Требуется считать данные о продолжительности периода, преобразовать их в тип Duration,
а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.

Sample Input:
12 мин. 13 сек.

Sample Output:
Fri May 15 19:28:18 UTC 2020
*/

func main() {
	const now int64 = 1589570165
	var inputTime string
	nr := bufio.NewReader(os.Stdin)
	inputTime, _ = nr.ReadString('\n')
	inputTime = strings.TrimRight(inputTime, "\n")
	inputTime = strings.TrimRight(inputTime, "\r")
	inputTime = strings.Replace(inputTime, " мин. ", "m", 1)
	inputTime = strings.Replace(inputTime, " сек.", "s", 1)

	parseDuration, err := time.ParseDuration(inputTime)
	if err != nil {
		panic(err)
	}
	unixTime := time.Unix(now, 0)
	fmt.Printf("%v\n", unixTime.UTC().Add(parseDuration).Format(time.UnixDate))
}
