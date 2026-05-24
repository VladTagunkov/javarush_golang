package main

import "fmt"

func main() {
	var (
		index    int
		found    bool
		foundVal int
		foundIdx int
	)

	for {
		var x int
		if _, err := fmt.Scan(&x); err != nil {
			break
		}
		if x == 0 {
			break
		}

		// TODO: Индекс должен увеличиваться только для чисел из диапазона 1..10000.

		if x <= 0 || x > 10000 {
			// TODO: Числа вне диапазона 1..10000 нужно игнорировать через continue и НЕ учитывать в нумерации.
			continue
		}
		index++

		// TODO: Проверьте, подходит ли число по условию (делится на 13 без остатка).
		// TODO: Если подходит — сохраните found/foundVal/foundIdx и немедленно завершите цикл через break.
		if x%13 == 0 {
			foundVal = x
			foundIdx = index
			found = true
			break
		}
	}

	if found {
		fmt.Printf("FOUND %d %d\n", foundVal, foundIdx)
		return
	}

	fmt.Println("NOT FOUND")
}
