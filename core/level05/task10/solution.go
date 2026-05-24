package main

import "fmt"

func main() {
	var playerName string
	fmt.Scan(&playerName)

	var count int

	// TODO: Обойдите строку playerName ровно одним циклом for _, v := range playerName (индекс должен быть проигнорирован через _).
	// TODO: Внутри цикла сравнивайте v с символом 'o' и увеличивайте count на 1 при каждом совпадении.
	// TODO: Не используйте готовые функции подсчёта (например, strings.Count), не добавляйте дополнительные циклы и коллекции для символов.
	for _, v := range playerName {
		if v == 'o' {
			count++
		}
	}

	fmt.Print(count)
}
