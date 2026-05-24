package main

import "fmt"

func main() {
	var city string
	var days int
	var perDay int

	// Ввод: одно слово + два целых числа.
	fmt.Scan(&city, &days, &perDay)

	total := 0
	// TODO: Посчитайте общую стоимость поездки в переменной total.
	total = days * perDay

	// TODO: Выведите ровно одну строку в требуемом формате (через fmt.Printf) и завершите её переводом строки \n.
	fmt.Printf("Trip to %s: %d days, $%d/day, total=$%d\n", city, days, perDay, total)
}
