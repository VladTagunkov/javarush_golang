package main

import (
	"fmt"
	"math"
)

func main() {
	var balance int
	var skippedOperations int

	for { // бесконечный цикл ввода по требованию
		var x int
		fmt.Scan(&x)

		if x == 0 { // 0 завершает ввод и не учитывается
			break
		}

		abs := math.Abs(float64(x))
		// TODO: Посчитайте модуль числа x и сохраните в переменную abs.

		if abs > 1000 {
			// TODO: Увеличьте skippedOperations и завершите текущую итерацию через continue.
			skippedOperations++
			continue
		}

		// TODO: Обновите balance для допустимой операции (когда |x| <= 1000).
		balance = balance + x
	}

	fmt.Printf("%d %d", balance, skippedOperations)
}
