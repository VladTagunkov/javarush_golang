package main

import (
	"fmt"
)

func main() {
	var operationsCount int
	fmt.Scan(&operationsCount)

	totalSum := 0

	if operationsCount < 0 {
		operationsCount = 0
	}

	for i := 0; i < operationsCount; i++ {
		var op int
		fmt.Scan(&op)

		totalSum += op
		// TODO: Добавьте op к totalSum (накапливайте сумму операций без хранения списка).
		// TODO: Убедитесь, что отрицательные значения учитываются корректно.
	}

	fmt.Println(totalSum)
}
