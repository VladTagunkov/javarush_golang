package main

import "fmt"

func main() {
	var systemTag string
	fmt.Scan(&systemTag)

	// Ответ по умолчанию: "не найдено" (паттерн index-or-minus-one)
	answer := -1

	for i, v := range systemTag {
		_ = i
		_ = v
		if v == '-' {
			answer = i
			break
		}
		// TODO: Найдите индекс первого символа '-' одним проходом по строке через range.
		// TODO: При первом обнаружении: присвойте answer индекс i и завершите цикл через break.
	}

	fmt.Println(answer)
}
