package main

import "fmt"

func main() {
	var trackingCode string
	fmt.Scan(&trackingCode)

	// Обходим строку только через range.
	for i := range trackingCode {
		// TODO: Выведите индекс i (каждый индекс на новой строке) без лишнего вывода.
		fmt.Println(i)
	}
}
