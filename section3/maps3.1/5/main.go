package main

/*
 * Группировка городов по численности населения в тысячах человек
 * от 10 до 100, от 100 до 1000 и более 1000:
 * groupCity map[int][]string{
 *	 10: []string{...},
 *	 100: []string{...},
 *	 1000: []string{...},
 * }
 *
 * Население городов в тысячах человек:
 * cityPopulation map[string]int{...}

 * Однако база данных с информацией о точной численности населения содержала ошибки, поэтому в cityPopulation
 * в т.ч. была сохранена информация о городах, которые входят в другие группы из groupCity.

 * Ваша программа имеет доступ к обоим указанным отображениям, требуется исправить cityPopulation,
 * чтобы в ней была сохранена информация только о городах из группы groupCity[100].
 */

func main() {
	var groupCity = make(map[int][]string)    //10,100,1000
	var cityPopulation = make(map[string]int) //cities

	for populationKey, _ := range cityPopulation {
		var isFound = false
		for _, city := range groupCity[100] {
			if city == populationKey {
				isFound = true
				break
			}
		}
		if !isFound {
			delete(cityPopulation, populationKey)
		}
	}
}
