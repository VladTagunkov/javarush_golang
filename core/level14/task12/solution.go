package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s) // читаем строку как один токен (без пробелов)

	r := []rune(s) // работаем с Unicode-символами, а не с байтами
	fmt.Println(len(r))

	// TODO: Разверните слайс рун r "на месте" обменами элементов (не более len(r)/2 обменов).
	// TODO: После разворота выведите результат строго через string(r).
	ran := len(r) / 2
	for i, _ := range r {
		if i < ran {
			r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
		}
	}

	s = string(r)
	fmt.Println(s)
}
