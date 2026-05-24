package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	mask := uint(0)
	// TODO: Вычислите маску включения ровно одного бита по номеру n (0..15) с помощью битового сдвига (тип результата должен быть uint).
	mask = uint(1) << uint(n)
	fmt.Printf("%d %016b", mask, mask)
}
