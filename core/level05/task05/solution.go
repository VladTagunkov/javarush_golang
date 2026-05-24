package main

import "fmt"

func main() {
	var currentPower int
	fmt.Scan(&currentPower)

	doubles := 0

	// for <cond> { ... } — это "for как while" (цикл с одним условием)
	for currentPower < 1000 {
		// TODO: Реализуйте изменение мощности (удвоение currentPower) в каждой итерации.
		// TODO: Реализуйте подсчёт количества удваиваний (увеличивайте doubles на 1 в каждой итерации).
		// TODO: Уберите/замените этот break, чтобы цикл действительно выполнялся до достижения порога.
		//break
		currentPower = currentPower * 2
		doubles++
	}

	// Формат: итоговое_значение количество_удваиваний
	fmt.Printf("%d %d", currentPower, doubles)
}
