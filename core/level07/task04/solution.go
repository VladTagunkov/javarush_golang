package main

import "fmt"

func main() {
	var capacitySlots uint
	var usedSlots uint

	fmt.Scan(&capacitySlots, &usedSlots)

	if usedSlots > capacitySlots {
		fmt.Println("error")
		var zero uint = 0
		fmt.Printf("%v (%T)\n", zero, zero)
	} else {
		// TODO: Добавьте проверку, что usedSlots > capacitySlots — это ошибка.
		// TODO: В ветке ошибки нельзя делать вычитание в uint (иначе будет underflow).
		// TODO: В ветке ошибки нужно вывести две строки: сначала "error", затем "0 (uint)" через fmt.Printf("%v (%T)\n", ...).
		freeSlots := capacitySlots - usedSlots

		// TODO: Если ошибки нет, нужно вывести ровно одну строку: freeSlots и его тип (uint) через fmt.Printf("%v (%T)\n", ...).
		fmt.Printf("%v (%T)\n", freeSlots, freeSlots)
	}
}
