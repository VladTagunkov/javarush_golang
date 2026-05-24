package main

import "fmt"

func main() {
	var firstMeasure float64
	var secondMeasure float64
	var thirdMeasure float64

	fmt.Scan(&firstMeasure, &secondMeasure, &thirdMeasure)

	var average float64 = 0
	// TODO: Вычислите среднее значение трёх измерений по формуле из условия.
	average = (firstMeasure + secondMeasure + thirdMeasure) / 3

	fmt.Printf("avg=%.2f\n", average)
	fmt.Printf("avg=%.17f\n", average)
}
