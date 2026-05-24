package main

import "fmt"

func main() {
	var startBalance, targetBalance, monthlyAdd int
	fmt.Scan(&startBalance, &targetBalance, &monthlyAdd)

	// Уже достигли цели — цикл не нужен.
	if startBalance >= targetBalance {
		fmt.Println(0)
		return
	}

	// Цель недостижима — важно обработать ДО входа в цикл.
	if monthlyAdd <= 0 {
		fmt.Println(-1)
		return
	}

	months := 0
	balance := startBalance

	// Цикл в стиле while: завершается только по условию (без break).
	for balance < targetBalance {
		// TODO: На каждой итерации увеличивайте balance на monthlyAdd (а не на 1).
		// TODO: На каждой итерации увеличивайте months на 1.
		balance = balance + monthlyAdd
		months++
	}

	fmt.Println(months)
}
