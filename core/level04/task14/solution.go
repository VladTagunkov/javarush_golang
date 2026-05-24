package main

import "fmt"

func main() {
	var password string
	fmt.Scan(&password)

	if len(password) == 0 {
		fmt.Print("empty")
		return
	}

	var lastIndex int
	var lastChar byte

	lastIndex = -1
	lastChar = '?'

	for i := 0; i < len(password); i++ {
		// TODO: Реализуйте обновление lastIndex и lastChar в цикле так, чтобы после завершения цикла
		// TODO: они содержали индекс и символ, соответствующие последнему символу строки password (без выхода за границы).
		lastIndex = i
		lastChar = password[i]
	}

	fmt.Printf("%d %c", lastIndex, lastChar)
}
