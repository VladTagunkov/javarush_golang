package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// По требованию: создаём слайс длины N через make и заполняем из ввода.
	tasks := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tasks[i])
	}

	var holeIndex int
	fmt.Scan(&holeIndex)

	// TODO: Проверьте, что holeIndex находится в диапазоне 0 ≤ holeIndex < N-1 (при N=0 или N=1 сдвиг не выполняется).
	// TODO: Если индекс корректный, выполните сдвиг элементов влево РОВНО одной операцией copy, не меняя длину tasks.
	if 0 <= holeIndex && holeIndex < n-1 {
		copy(tasks[holeIndex:], tasks[holeIndex+1:])
	}

	// Ровно N чисел через пробел.
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(tasks[i])
	}
	fmt.Println()
}
