package main

import "fmt"

func main() {
	var usedSeconds, pricePerMinuteCents, baseFeeCents int
	fmt.Scan(&usedSeconds, &pricePerMinuteCents, &baseFeeCents)

	billableMinutes := usedSeconds / 60
	leftoverSeconds := usedSeconds % 60
	// TODO: Посчитайте количество полных оплачиваемых минут (billableMinutes) из usedSeconds.
	// TODO: Посчитайте остаток секунд (leftoverSeconds), который не попал в полные минуты.

	variableCents := billableMinutes * pricePerMinuteCents
	totalCents := baseFeeCents + variableCents
	// TODO: Посчитайте переменную часть стоимости (variableCents) в центах.
	// TODO: Посчитайте итоговую стоимость (totalCents) в центах с учётом базовой платы.

	euros := totalCents / 100
	cents := totalCents % 100
	// TODO: Разбейте totalCents на евро (euros) и центы (cents) с помощью целочисленной арифметики.

	fmt.Printf("%d %02d %d", euros, cents, leftoverSeconds)
}
