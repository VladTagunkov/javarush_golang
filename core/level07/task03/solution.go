package main

import "fmt"

func main() {
	// Все суммы храним в центах, поэтому используем int64.
	var incomeCents int64
	var expenseCents int64
	fmt.Scan(&incomeCents, &expenseCents)

	// Баланс тоже строго int64 (никаких int/float).
	var balanceCents int64
	balanceCents = incomeCents - expenseCents
	// TODO: вычислите balanceCents как incomeCents - expenseCents (все переменные int64)

	// Первая строка "квитанции": значение баланса + его тип через %T.
	fmt.Printf("balanceCents=%d type=%T\n", balanceCents, balanceCents)
	// TODO: выведите тип именно balanceCents через %T (не используйте литералы)

	// Вторая строка: статус по знаку баланса (строго одно слово).
	if balanceCents == 0 {
		fmt.Println("zero")
	} else if balanceCents < 0 {
		fmt.Println("loss")
	} else if balanceCents > 0 {
		fmt.Println("profit")
	}
	//fmt.Println("status")
	// TODO: определите статус по balanceCents через if/else if/else и выведите profit/loss/zero
}
