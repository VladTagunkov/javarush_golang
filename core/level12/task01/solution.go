package main

import "fmt"

func main() {
	// Ровно один заданный исходный слайс по условию.
	s := []int{3, 1, 4, 1, 5, 9}

	var n int
	fmt.Scan(&n)
	if n < 0 || n > len(s) {
		fmt.Println([]int{})
	} else {
		fmt.Println(s[:n])
	}

	// TODO: Проверьте, что n находится в диапазоне [0..len(s)] и программа не падает при любом вводе.
	// TODO: Если n некорректное (n < 0 или n > len(s)) — выведите пустой слайс [] через fmt.Println([]int{}).
	// TODO: Если n корректное — выведите первые n элементов строго через нарезку s[:n] (без циклов и без s[i]).

	//fmt.Println([]int{})
}
