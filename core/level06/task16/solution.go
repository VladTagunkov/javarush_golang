package main

import "fmt"

func main() {
	var item string
	var qty, price int

	fmt.Scan(&item, &qty, &price)

	var sum int
	sum = qty * price
	// TODO: Вычислите сумму покупки: sum = qty * price.

	// TODO: Выведите первую строку чека в нужном формате через fmt.Printf (используйте %s для item и %d для целых).
	fmt.Printf("Item=%s Qty=%d Price=%d\n", item, qty, price)
	fmt.Printf("Sum=%d\n", sum)
	fmt.Printf("DEBUG: itemType=%T qtyType=%T priceType=%T sumType=%T\n", item, qty, price, sum)
	// TODO: Выведите вторую строку с суммой в нужном формате через fmt.Printf (используйте %d).
	// TODO: Выведите третью строку DEBUG с типами переменных item, qty, price, sum через fmt.Printf (используйте %T).
}
