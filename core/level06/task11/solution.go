package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		// По условию: ничего не читаем дальше
		fmt.Println(0)
		return
	}

	// По требованию: объявляем через var и используем zero values (0 и false)
	var max int
	var hasMax bool

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if i == 0 {
			max = x
			hasMax = true
		}

		// TODO: Реализуйте корректное обновление максимума с учётом отрицательных чисел.
		// TODO: Используйте hasMax (zero value = false), чтобы на первом считанном значении
		// TODO: установить max, а на остальных — сравнивать с текущим max.
		if !hasMax {
			hasMax = true
		}
		if x > max {
			max = x
		}
	}

	fmt.Println(max)
}
