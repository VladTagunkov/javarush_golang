package main

import "fmt"

const minAge = 18       // untyped-константа
const maxLen uint8 = 12 // typed-константа

func main() {
	var nickname string
	var age int
	fmt.Scan(&nickname, &age)

	allowed := false

	if (age >= minAge) && (len(nickname) < int(maxLen)) {
		allowed = true
	} else {
		allowed = false
	}
	// TODO: Реализуйте условие допуска: возраст должен быть не меньше minAge, а длина ника (len) не больше maxLen.
	// TODO: В сравнении len(nickname) и maxLen сделайте явное приведение типов.
	// TODO: В логике допуска используйте && (должны выполниться оба условия одновременно).

	if allowed {
		fmt.Println("ALLOWED")
		return
	}
	fmt.Println("DENIED")
}
