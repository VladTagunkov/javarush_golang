package main

import "fmt"

func main() {
	var n int
	var text string
	fmt.Scan(&n, &text)

	// По условию: при N <= 0 печатаем пустую строку.
	if n <= 0 {
		fmt.Println()
		return
	}

	// TODO: Реализуйте обрезку по первым N Unicode-рунам, а не по байтам.
	// TODO: Нельзя преобразовывать строку в []rune.
	// TODO: Используйте for range по строке text, чтобы найти байтовый индекс начала (N+1)-й руны.
	// TODO: Если в строке меньше либо ровно N рун — выведите строку целиком без изменений.
	count := 0
	cut := 0
	for i := range text {
		if count == n {
			cut = i
			break
		}
		count++
	}
	// Временная упрощённая реализация: режем по байтам (это неверно для UTF-8).
	//cut := len(text)
	if cut == 0 {
		fmt.Println()
		return
	}
	if count == n {
		fmt.Println(text)
		return
	}

	if n < len(text) {
		cut = n
	}

	fmt.Println(text[:cut])
}
