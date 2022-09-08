package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:
2020-05-15 08:00:00

Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату на стандартный вывод в том же формате.
Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день,
а затем вывести на стандартный вывод в том же формате.
*/

func main() {
	var inputTime string
	nr := bufio.NewReader(os.Stdin)
	inputTime, _ = nr.ReadString('\n')
	inputTime = strings.Trim(inputTime, "\n")
	inputTime = strings.Trim(inputTime, "\r")
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, inputTime)
	if err != nil {
		fmt.Println(err)
	}
	if parsedTime.Hour() >= 13 {
		parsedTime = parsedTime.AddDate(0, 0, 1)
		fmt.Print(parsedTime.Format(layout))
	} else {
		fmt.Print(parsedTime.Format(layout))
	}
}
