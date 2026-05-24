package main

import "fmt"

func main() {
	// Справочник задан заранее (не вводится пользователем).
	ages := map[string]int{
		// TODO: добавьте несколько пар (имя → возраст); обязательно добавьте хотя бы одну запись с возрастом 0
		"alex":  33,
		"bob":   44,
		"piter": 0,
	}

	var personName string
	fmt.Scan(&personName)

	// TODO: найдите возраст по имени в map через ok-идиому, чтобы различать:
	// TODO: 1) ключа нет (ok == false)
	// TODO: 2) ключ есть, но возраст может быть 0 (ok == true, age == 0)
	// TODO: в зависимости от ok выведите либо "age: <число>", либо "no such person" (точный формат)
	if age, exists := ages[personName]; exists {
		fmt.Printf("age: %d\n", age)
	} else {
		fmt.Println("no such person")
	}

}
