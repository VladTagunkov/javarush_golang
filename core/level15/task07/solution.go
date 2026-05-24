package main

import "fmt"

func main() {
	// Предзаданный справочник скидок (важно: 0 — валидная скидка, не "нет данных")
	discounts := map[string]int{
		"alice": 10,
		"bob":   0,
		"carol": 5,
		"dave":  15,
	}

	var n int
	fmt.Scan(&n)

	missing := 0
	for i := 0; i < n; i++ {
		var login string
		fmt.Scan(&login)

		// TODO: Реализуйте корректную проверку наличия логина в map через ok-идиому (v, ok := discounts[login]).
		// TODO: Учтите, что скидка 0 — это валидное значение (если ok == true, нужно печатать discount: 0).
		// TODO: Увеличивайте missing только если логин отсутствует (ok == false).
		//v := discounts[login]
		if v, ok := discounts[login]; ok {
			fmt.Printf("login: %s discount: %d\n", login, v)
		} else {
			fmt.Printf("login: %s no discount data\n", login)
			missing += 1
		}

		// Текущая логика намеренно неполная: она путает "нет данных" и "скидка = 0".
		//if v == 0 {
		//	fmt.Printf("login: %s no discount data\n", login)
		//	missing++
		//	continue
		//}

		//fmt.Printf("login: %s discount: %d\n", login, v)
	}

	fmt.Printf("missing: %d\n", missing)
}
