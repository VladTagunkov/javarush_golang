package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// Фиксированная длина: заполняем по индексам, без append.
	scores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
	}

	var m int
	fmt.Scan(&m)

	applied := 0
	for i := 0; i < m; i++ {
		var idx, value int
		fmt.Scan(&idx, &value)

		// TODO: Реализуйте безопасное применение обновления scores[idx] = value:
		// TODO: - сначала проверьте, что 0 <= idx && idx < len(scores)
		// TODO: - только при корректном idx делайте присваивание scores[idx] = value
		// TODO: - увеличивайте applied только для успешно применённых обновлений
		if 0 <= idx && idx < len(scores) {
			scores[idx] = value
			applied++
		}
		// Сейчас обновления намеренно игнорируются (шаблон для студента).
	}

	// Ровно две строки вывода по условию.
	fmt.Println(applied)
	fmt.Println(scores)
}
