package main

import "fmt"

func main() {
	var command string
	fmt.Scan(&command)

	// TODO: Обойдите строку command циклом range только по индексу (одна переменная в range).
	// TODO: На каждой итерации печатайте текущее значение индекса i в отдельной строке через fmt.Println.
	// TODO: Убедитесь, что вывод соответствует реальным индексам, которые даёт range по строке в Go.
	for i := range command {
		fmt.Println(i)
	}
}
