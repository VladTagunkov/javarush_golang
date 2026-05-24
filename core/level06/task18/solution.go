package main

import "fmt"

func main() {
	var from, to int
	fmt.Scan(&from, &to)

	if from > to {
		// TODO: Если границы введены в обратном порядке, поменяйте значения from и to местами
		// TODO: Используйте множественное присваивание и оператор "=", чтобы не создать новые переменные (без shadowing)
		from, to = to, from
	}

	fmt.Printf("range: %d..%d\n", from, to)
}
