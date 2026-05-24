package main

import "fmt"

func main() {
	var targetValue int
	fmt.Scan(&targetValue) // маркер читаем один раз до обработки ленты

	pos := 1
	answer := -1
	found := false

	for {
		var x int
		if _, err := fmt.Scan(&x); err != nil {
			break // читаем, пока читается (EOF/ошибка завершает цикл)
		}

		//_ = x // технически используем переменную, пока логика не реализована
		if x == targetValue {
			found = true
			answer = pos
			break
		}
		pos++
		// TODO: Реализуйте проверку: если текущее число равно targetValue — это первое совпадение.
		// TODO: При первом совпадении сохраните позицию (pos) в answer, установите found = true и выйдите из цикла (break).
		// TODO: Корректно обновляйте pos: позиция считается среди чисел ПОСЛЕ targetValue, начиная с 1.

		// сейчас увеличивается всегда — это намеренно неполная логика, студент должен исправить
	}

	if !found {
		fmt.Println(-1)
		return
	}
	fmt.Println(answer)
}
