package main

import "fmt"

func main() {
	// Чтение множителя из stdin через fmt.Scan.
	var maxMultiplier int
	fmt.Scan(&maxMultiplier)

	// TODO: Реализуйте корректные границы циклов:
	// TODO: внешний цикл должен перебирать i от 1 до maxMultiplier включительно,
	// TODO: внутренний цикл должен перебирать j от 1 до maxMultiplier включительно.
	for i := 1; i <= maxMultiplier; i++ {
		for j := 1; j <= maxMultiplier; j++ {
			// TODO: Посчитайте произведение i*j и сохраните в product.
			product := i * j

			// TODO: Выведите строку строго в формате i*j=product (без пробелов), каждая строка с новой строки.
			fmt.Printf("%d*%d=%d\n", i, j, product)
		}
	}
}
