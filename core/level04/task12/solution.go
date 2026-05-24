package main

import "fmt"

func main() {
	var levelRow string
	fmt.Scan(&levelRow)

	var fixedRow string
	var replacedCount int

	for _, v := range levelRow {
		if v == '#' {
			fixedRow += "*"
			replacedCount++
		} else {
			fixedRow += string(v)
		}
		// TODO: Реализуйте обработку каждого символа строки через range.
		// TODO: Если текущий символ равен '#', добавьте в fixedRow "*" и увеличьте replacedCount.
		// TODO: Иначе добавьте в fixedRow текущий символ без изменений (как строку).
	}

	fmt.Println(fixedRow)
	fmt.Println(replacedCount)
}
