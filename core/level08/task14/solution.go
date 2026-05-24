package main

import (
	"fmt"
)

func main() {
	var flags, k uint
	fmt.Scan(&flags, &k)

	var mask uint
	mask = 0
	// TODO: Вычислите маску для бита k с помощью битового сдвига (нужен беззнаковый тип)
	mask = 1 << k
	// TODO: Определите, был ли бит k установлен ДО изменения (проверка через побитовое И),
	// TODO: и выведите "present" или "absent" ровно в одной строке.
	if (flags & mask) != 0 {
		fmt.Println("present")
	} else {
		fmt.Println("absent")
	}
	newFlags := flags | mask
	// TODO: Сформируйте новое значение флагов так, чтобы бит k был включён (операция побитового ИЛИ)

	fmt.Printf("%d %08b\n", newFlags, newFlags)
}
