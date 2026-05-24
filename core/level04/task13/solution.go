package main

import "fmt"

func main() {
	var ticketsCount int
	fmt.Scan(&ticketsCount)

	// Пустая очередь: печатаем ровно одну строку и сразу завершаемся.
	if ticketsCount <= 0 {
		fmt.Println("empty")
		return
	}

	// TODO: Реализуйте вывод ровно ticketsCount талонов, начиная с 0 (каждый номер — с новой строки).
	// TODO: Проверьте границы цикла, чтобы не было лишнего значения и не было пропусков.
	for i := 0; i < ticketsCount; i++ {
		fmt.Println(i)
		// TODO: Напечатайте номер талона для текущей итерации.
	}
}
