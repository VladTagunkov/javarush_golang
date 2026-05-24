package main

import "fmt"

func main() {
	var (
		x     int
		count int
		sum   int
	)

	for { // потенциально бесконечный цикл чтения
		if _, err := fmt.Scan(&x); err != nil { // 1 число за итерацию; ошибка (например EOF) — выход
			break
		}

		if x == 0 {
			break
		}

		if x < 0 {
			// TODO: Отрицательные значения нужно игнорировать с помощью continue (не менять статистику).
			continue
		}

		if x > 0 {
			// TODO: Учитывайте только положительные значения: увеличьте count и добавьте x к sum.
			count++
			sum += x
		}
	}

	fmt.Printf("count=%d sum=%d", count, sum)
}
