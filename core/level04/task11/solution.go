package main

import "fmt"

func main() {
	var slug string
	fmt.Scan(&slug)

	pos := -1
	for i, v := range slug {
		if v == '-' {
			pos = i
			break
		}
	}
	// TODO: Найдите индекс первого символа '-' в строке slug.
	// TODO: Обход должен быть через: for i, v := range slug
	// TODO: При первом совпадении (v == '-') сохраните индекс i в pos и сразу выйдите из цикла через break.
	// TODO: Если дефиса нет — pos должен остаться равным -1.

	fmt.Print(pos)
}
