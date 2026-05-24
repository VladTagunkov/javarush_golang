package main

import "fmt"

func main() {
	var totalUnits, partsCount int
	fmt.Scan(&totalUnits, &partsCount)

	integerQuotient := totalUnits / partsCount

	// TODO: Исправьте вычисление дробного частного: сейчас дробная часть теряется из-за целочисленного деления.
	floatQuotient := float64(totalUnits) / float64(partsCount)

	fmt.Println(integerQuotient)
	fmt.Println(floatQuotient)
}
