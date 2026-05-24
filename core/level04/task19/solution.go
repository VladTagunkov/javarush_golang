package main

import "fmt"

func main() {
	var participantsCount int
	fmt.Scan(&participantsCount)

	// Ранний выход: если участников нет, не читаем результаты.
	if participantsCount <= 0 {
		fmt.Println(0)
		fmt.Println(-1)
		return
	}

	// Считываем первое значение и используем его как стартовое.
	var maxValue int
	fmt.Scan(&maxValue)
	maxPos := 1

	// Читаем оставшиеся значения в цикле, без хранения всей последовательности.
	for i := 2; i <= participantsCount; i++ {
		var v int
		fmt.Scan(&v)

		if v > maxValue {
			// TODO: Исправьте обновление максимума и позиции так, чтобы позиция соответствовала первому вхождению максимального значения.
			maxValue = v
			maxPos = i
		}
	}

	fmt.Println(maxValue)
	fmt.Println(maxPos)
}
