package main

import "fmt"

/**
 * Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно.
 * У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.
 * Если значение On == false, то оба метода вернут false.
 * Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true),
 * если его нет, то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.
 * Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой структуры
 * с именем testStruct в функции main, в дальнейшем программа проверит результат.
 */

type TestStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (play *TestStruct) Shoot() bool {
	if play.On == false {
		return false
	}
	if play.Ammo > 0 {
		play.Ammo--
		return true
	}
	return false
}

func (play *TestStruct) RideBike() bool {
	if play.On == false {
		return false
	}
	if play.Power > 0 {
		play.Power--
		return true
	}
	return false
}

func main() {

	testStruct := &TestStruct{
		On:    true,
		Ammo:  1,
		Power: 2,
	}

	testStruct.Shoot()
	fmt.Println(testStruct.Ammo)
	testStruct.RideBike()
	fmt.Println(testStruct.Power)

}
