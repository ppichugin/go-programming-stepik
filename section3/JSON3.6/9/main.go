package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
Данная задача ориентирована на реальную работу с данными в формате JSON.
Для решения мы будем использовать справочник ОКВЭД (Общероссийский классификатор видов экономической деятельности),
который был размещен на web-портале data.gov.ru.

Необходимая вам информация о структуре данных содержится в файле structure-20190514T0000.json, а сами данные,
которые вам потребуется декодировать, содержатся в файле data-20190514T0100.json. Файлы размещены в нашем репозитории
на github.com: https://github.com/semyon-dev/stepik-go/tree/master/work_with_json

Для того, чтобы показать, что вы действительно смогли декодировать документ вам необходимо в качестве ответа
записать сумму полей global_id всех элементов, закодированных в файле.
*/

func main() {
	const URL = "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	target := StructClassify{}
	json.NewDecoder(resp.Body).Decode(&target)
	var sum uint64
	for _, elem := range target {
		sum += elem.GlobalId
	}
	fmt.Println(sum)
}

type StructClassify []struct {
	GlobalId uint64 `json:"global_id"`
}
