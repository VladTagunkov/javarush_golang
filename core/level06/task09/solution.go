package main

import "fmt"

func main() {
	// TODO: Объявите четыре переменные через var без явной инициализации: int, float64, bool и string.
	var i int
	var f float64
	var b bool
	var s string

	// TODO: Выведите значения ровно в 4 строках и в строгом порядке: int, float, bool, string.
	// TODO: Используйте точные префиксы: "int:", "float:", "bool:", "string:".
	// TODO: Строковое значение выведите обрамлённым маркерами '>' и '<', чтобы пустая строка выглядела как ><.

	// Временный вывод, чтобы код оставался запускаемым, но не соответствовал требуемому формату.
	//fmt.Println(i, f, b, s)
	fmt.Printf("int: %d\n", i)
	fmt.Printf("float: %g\n", f)
	fmt.Printf("bool: %t\n", b)
	fmt.Printf("string: >%s<\n", s)
}
