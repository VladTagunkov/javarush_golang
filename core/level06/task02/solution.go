package main

import "fmt"

func main() {
	total := 0
	fmt.Scan(&total)

	discount := 0
	// TODO: Реализуйте правило расчёта скидки по условию задачи (discount должен быть доступен после if).
	if total >= 1000 {
		// TODO: Назначьте discount корректное значение через обычное присваивание (=), без переобъявления (:=).
		discount = 100
	} else {
		// TODO: Назначьте discount корректное значение через обычное присваивание (=), без переобъявления (:=).
		discount = 0
	}

	// TODO: Вычислите finalTotal после if как total - discount (отдельной переменной).
	finalTotal := total - discount

	fmt.Printf("%d %d\n", discount, finalTotal)
}
