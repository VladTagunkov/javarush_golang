package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	doubled := 2 * n
	// TODO: Вычислите удвоенное значение числа n и запишите его в doubled.

	// Важно: не называем строковую переменную "fmt", иначе затенится пакет fmt.
	outFormat := "%d %d\n"
	fmt.Printf(outFormat, n, doubled)
}
