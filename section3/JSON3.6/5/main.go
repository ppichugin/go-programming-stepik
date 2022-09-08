package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
На стандартный ввод подаются данные о студентах университетской группы в формате JSON.
В сведениях о каждом студенте содержится информация о полученных им оценках (Rating).
Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы.
Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:
{
    "Average": 14.1
}
*/

func main() {
	data, _ := ioutil.ReadAll(os.Stdin)
	var myGroup Group
	if err := json.Unmarshal(data, &myGroup); err != nil {
		fmt.Println(err)
		return
	}
	var students, marks int
	for i := 0; i < len(myGroup.Students); i++ {
		students++
		for j := 0; j < len(myGroup.Students[i].Rating); j++ {
			marks++
		}
	}
	fmt.Println("Students: ", students)
	res := Result{float32(marks) / float32(students)}
	avg, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(avg))
}

type Result struct {
	Average float32 `json:"Average"`
}

type Group struct {
	ID       int    `json:"ID"`
	Number   string `json:"Number"`
	Year     int    `json:"Year"`
	Students []struct {
		LastName   string `json:"LastName"`
		FirstName  string `json:"FirstName"`
		MiddleName string `json:"MiddleName"`
		Birthday   string `json:"Birthday"`
		Address    string `json:"Address"`
		Phone      string `json:"Phone"`
		Rating     []int  `json:"Rating"`
	} `json:"Students"`
}
