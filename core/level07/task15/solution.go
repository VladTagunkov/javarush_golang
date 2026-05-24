package main

import "fmt"

func main() {
	var incomingValue int64
	fitsInt8 := false
	fmt.Scan(&incomingValue)

	if (incomingValue <= 127) && (incomingValue >= -128) {
		//fmt.Println(int8(incomingValue))
		fitsInt8 = true
	} else {
		//fmt.Println("OVERFLOW")
		fitsInt8 = false
	}
	// TODO: Реализуйте явную проверку, помещается ли incomingValue в диапазон int8.
	// TODO: Используйте if/else: в ветке успеха — безопасно преобразуйте к int8 и выведите число,
	// TODO: иначе — выведите строку OVERFLOW.

	if fitsInt8 {
		// TODO: Выполните преобразование к int8 только после успешной проверки диапазона.
		fmt.Println(int8(incomingValue))
	} else {
		fmt.Println("OVERFLOW")
	}
}
