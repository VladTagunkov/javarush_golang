package main

import "fmt"

func main() {
	var ordersCount int
	fmt.Scan(&ordersCount)

	totalRevenue := 0

	for i := 0; i < ordersCount; i++ {
		var orderTotal int
		fmt.Scan(&orderTotal)

		// TODO: Реализуйте накопление суммы всех заказов в переменной totalRevenue
		// Сейчас логика неверная: сохраняется только значение текущего заказа.
		totalRevenue += orderTotal
	}

	fmt.Println(totalRevenue)
}
