package main

import "fmt"

func main() {
	var login string
	fmt.Scan(&login)

	invalid := -1

	for i, v := range login {
		if i == 0 {
			if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v == '_') {
				invalid = -1
			} else {
				invalid = i
				break
			}
			// TODO: Реализовать проверку первого символа логина по правилам задачи.
			// TODO: Если символ недопустим — установить invalid = i и завершить цикл через break.
			continue
		}

		_ = v
		if i != 0 {
			if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') || (v == '_') {
				invalid = -1
			} else {
				invalid = i
				break
			}
		}
		// TODO: Реализовать проверку остальных символов логина по правилам задачи.
		// TODO: Если символ недопустим — установить invalid = i и завершить цикл через break.
	}

	if invalid == -1 {
		fmt.Print("OK")
	} else {
		fmt.Printf("INVALID %d", invalid)
	}
}
