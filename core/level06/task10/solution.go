package main

import "fmt"

func main() {
	// zero value у bool = false, поэтому found не инициализируем явно
	var found bool

	for i := 0; i < 5; i++ {
		var x int
		fmt.Scan(&x)
		if x == 10 {
			found = true
		}

		// TODO: Реализуйте проверку: если среди вводимых чисел встречается 10,
		// TODO: то переменная found должна стать true (достаточно одного совпадения).
	}

	// Печатаем ровно одно слово: true или false
	fmt.Println(found)
}
