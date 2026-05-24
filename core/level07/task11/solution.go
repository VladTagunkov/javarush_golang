package main

import "fmt"

func main() {
	var unitPriceCents, quantity, discountPercent int
	fmt.Scan(&unitPriceCents, &quantity, &discountPercent)

	// Все считаем в центах, без float.
	subtotalCents := unitPriceCents * quantity

	// TODO: Посчитайте discountCents (скидку в центах) строго целочисленными операциями.
	// TODO: Важно сохранить результат в отдельной переменной discountCents.
	discountCents := (subtotalCents * discountPercent) / 100

	// TODO: Посчитайте totalCents через subtotalCents и discountCents (отдельной переменной).
	totalCents := subtotalCents - discountCents

	eur := totalCents / 100
	cen := totalCents % 100

	fmt.Printf("%d %02d", eur, cen)
}
