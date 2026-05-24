package main

import "fmt"

func main() {
	var baseCost, extraUnits, multiplier int64
	fmt.Scan(&baseCost, &extraUnits, &multiplier)

	var estimateA int64
	var estimateB int64

	// TODO: Вычислите две оценки стоимости по условию задачи и присвойте их переменным estimateA и estimateB.
	estimateA = baseCost + extraUnits*multiplier
	estimateB = (baseCost + extraUnits) * multiplier

	fmt.Print(estimateA, " ", estimateB)
}
