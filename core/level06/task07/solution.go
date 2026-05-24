package main

import "fmt"

func main() {
	var x, y, z float64
	var minVal float64
	var maxVal float64
	fmt.Scan(&x, &y, &z)

	// TODO: Найдите минимальное из трёх значений x, y, z (тип float64) только через if/else
	if x <= y && x <= z {
		minVal = x
	} else if y <= x && y <= z {
		minVal = y
	} else if z <= x && z <= y {
		minVal = z
	}

	// TODO: Найдите максимальное из трёх значений x, y, z (тип float64) только через if/else
	if x >= y && x >= z {
		maxVal = x
	} else if y >= x && y >= z {
		maxVal = y
	} else if z >= x && z >= y {
		maxVal = z
	}

	avg := (x + y + z) / 3

	// TODO: Вычислите stable по правилу: (maxVal - minVal) <= 1.0.
	stable := (maxVal - minVal) <= 1.0

	fmt.Printf("min=%v\n", minVal)
	fmt.Printf("max=%v\n", maxVal)
	fmt.Printf("avg=%v\n", avg)
	fmt.Printf("stable=%v\n", stable)
}
