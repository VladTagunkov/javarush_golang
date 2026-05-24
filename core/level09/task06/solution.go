package main

import "fmt"

// minMax3 возвращает сначала минимум, затем максимум из трёх чисел.
func minMax3(a, b, c int) (int, int) {
	// TODO: Реализуйте поиск минимума и максимума среди a, b, c.
	// TODO: Используйте только сравнения (<, >) и ветвления if/else (без циклов, массивов и слайсов).
	// TODO: Важно: сначала верните min, затем max (порядок результатов важен).
	mn := a
	mx := b
	if (a >= b) && (a >= c) {
		mx = a
	} else if (b >= c) && (b >= a) {
		mx = b
	} else if (c >= a) && (c >= b) {
		mx = c
	}
	if (a <= b) && (a <= c) {
		mn = a
	} else if (b <= c) && (b <= a) {
		mn = b
	} else if (c <= a) && (c <= b) {
		mn = c
	}

	// Временная заглушка: возвращает значения, которые не являются корректными для общего случая.
	return mn, mx
}

func main() {
	var measureA, measureB, measureC int
	fmt.Scan(&measureA, &measureB, &measureC)

	mn, mx := minMax3(measureA, measureB, measureC)

	fmt.Printf("min=%d\n", mn)
	fmt.Printf("max=%d\n", mx)
}
