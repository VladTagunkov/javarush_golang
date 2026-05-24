package main

import "fmt"

func main() {
	var levelsCount int
	fmt.Scan(&levelsCount)

	// Отрицательное значение — ошибка входных данных.
	if levelsCount < 0 {
		fmt.Print("error")
		return
	}

	var totalRewards int
	for level := 1; level <= levelsCount; level++ {
		// TODO: Реализуйте накопление суммы 1 + 2 + … + levelsCount в переменной totalRewards.
		totalRewards += level
	}

	// Печатаем сумму и тип переменной (должно быть int).
	fmt.Printf("%d (%T)", totalRewards, totalRewards)
}
