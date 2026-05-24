package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var title string
		fmt.Scan(&title)
		// TODO: Считайте очередное название задачи в переменную title с помощью fmt.Scan(&title)

		// TODO: Выведите строку результата строго в формате "<номер> - <название>\n" без лишнего вывода
		fmt.Printf("%d - %s\n", i, title) // временный вывод, намеренно не соответствует требованиям
	}
}
