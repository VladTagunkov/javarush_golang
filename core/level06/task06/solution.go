package main

import "fmt"

func main() {
	var login string
	var code int

	fmt.Scan(&login, &code)

	var isAdmin bool = login == "admin"
	// TODO: Сравните login с "admin" и запишите результат в isAdmin (только через ==).

	var isCodeOK bool = code == 4321
	// TODO: Сравните code с 4321 и запишите результат в isCodeOK (только через ==).

	var canEnter bool = isAdmin && isCodeOK
	// TODO: Вычислите canEnter как isAdmin && isCodeOK (только через &&).

	fmt.Printf("canEnter=%v\n", canEnter)

	if canEnter {
		fmt.Println("ACCESS GRANTED")
	} else {
		fmt.Println("ACCESS DENIED")
	}
}
