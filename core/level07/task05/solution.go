package main

import "fmt"

func main() {
	// TODO: Создайте переменную defaultReading через := и присвойте ей значение 0.1 (без явного указания типа).
	defaultReading := 0.1
	// TODO: Создайте переменную liteReading с явным типом float32 и присвойте ей значение 0.1.
	var liteReading float32
	liteReading = 0.1

	fmt.Printf("%f %T\n", defaultReading, defaultReading)
	fmt.Printf("%f %T\n", liteReading, liteReading)
	// TODO: Выведите ровно две строки через fmt.Printf: по одной строке для каждой переменной.
	// TODO: Формат обеих строк должен быть одинаковым и включать одновременно значение переменной и её тип (тип выводите через %T), каждая строка должна завершаться переводом строки.
}
