package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// Важно: объявлены через var без явной инициализации (zero values).
	var sum float64   // сумма положительных
	var count float64 // количество положительных (тоже float64 по условию)
	var hasAny bool   // был ли хотя бы один положительный

	for i := 0; i < n; i++ {
		var x float64
		fmt.Scan(&x)
		if x > 0 {
			hasAny = true
			sum += x
			count = count + 1.0
		}

		// TODO: Реализуйте отбор только строго положительных значений (> 0).
		// TODO: Для положительных значений обновляйте sum, count (на 1.0) и hasAny.
	}

	// TODO: Реализуйте вывод: если положительных нет — вывести 0, иначе вывести sum / count (делить только если count != 0).
	if !hasAny || count == 0 {
		fmt.Println(0)
		return
	} else {
		if count != 0 {
			fmt.Println(sum / count)
		}
	}

	//fmt.Println(sum / count)
}
