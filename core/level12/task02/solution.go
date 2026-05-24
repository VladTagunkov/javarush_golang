package main

import "fmt"

func main() {
	// Фиксированный исходный слайс по условию (не модифицируем).
	s := []int{10, 20, 30, 40, 50, 60}

	var i, n int
	fmt.Scan(&i, &n)

	valid := false
	if (i >= 0 && n >= 0 && i+n <= len(s)) || (i == len(s) && n == 0) {
		valid = true
	}
	// TODO: Реализуйте проверку, что окно с параметрами i и n корректно для слайса s
	// TODO: При некорректных параметрах нужно вывести пустой слайс и завершиться без panic

	if !valid {
		fmt.Println([]int{})
		return
	}

	left, right := i, i+n
	// TODO: Вычислите границы окна (left/right) по введённым i и n
	// TODO: Сформируйте результат только через slicing исходного слайса s

	fmt.Println(s[left:right])
}
