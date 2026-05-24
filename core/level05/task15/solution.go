package main

import "fmt"

func main() {
	var partCode string
	fmt.Scan(&partCode)

	for i := 0; i < len(partCode); i++ {
		fmt.Printf("%d:%c\n", i, partCode[i])
	}

	// TODO: Напечатайте “паспорт” строки partCode: для каждого индекса i (от 0 до len(partCode)-1)
	// TODO: выведите строку в формате i:c (индекс, двоеточие, символ) на отдельной строке.
	// TODO: Используйте классический цикл for с условием строго вида i < len(partCode) и инкрементом i++.
	// TODO: Внутри цикла получайте символ только через partCode[i] и не печатайте ничего для пустой строки.
}
