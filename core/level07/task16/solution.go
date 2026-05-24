package main

import "fmt"

func main() {
	var firstByteValue, secondByteValue int
	fmt.Scan(&firstByteValue, &secondByteValue)

	outOfRange := false
	if ((firstByteValue >= 0) && (firstByteValue <= 255)) && ((secondByteValue >= 0) && (secondByteValue <= 255)) {
		outOfRange = false
	} else {
		outOfRange = true
	}
	// TODO: Реализуйте проверку, что оба числа в диапазоне 0..255 (включительно).
	// TODO: Если хотя бы одно число вне диапазона — установите outOfRange = true.

	if outOfRange {
		fmt.Print("OUT")
		return
	}

	var firstByte, secondByte uint8
	firstByte = uint8(firstByteValue)
	secondByte = uint8(secondByteValue)
	// TODO: Явно преобразуйте входные значения из int в uint8 и присвойте в firstByte и secondByte.

	var wrappedSum uint8
	// TODO: Посчитайте wrappedSum как сумму двух uint8 так, чтобы результат оставался uint8 (с переполнением по кругу).
	wrappedSum = firstByte + secondByte

	var exactSum int
	// TODO: Посчитайте exactSum как сумму в int, явно расширив каждое слагаемое uint8 -> int перед сложением.
	exactSum = int(firstByte) + int(secondByte)

	fmt.Printf("%d %d", wrappedSum, exactSum)
}
