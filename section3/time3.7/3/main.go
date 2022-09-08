package main

import (
	"fmt"
	"time"
)

/*
На стандартный ввод подается строковое представление даты и времени в следующем формате:
1986-04-16T05:20:00+06:00

Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:
Wed Apr 16 05:20:00 +0600 1986
*/

func main() {
	var dateStr string
	fmt.Scanln(&dateStr)

	parsedDateTime, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedDateTime.Format(time.UnixDate))
}
