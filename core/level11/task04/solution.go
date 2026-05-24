package main

import "fmt"

func EqualInts(a, b []int) bool {
	// Если длины разные — дальше сравнивать бессмысленно
	if len(a) != len(b) {
		return false
	}

	// TODO: Реализуйте поэлементное сравнение слайсов по индексам.
	// TODO: Нельзя сравнивать слайсы целиком через ==, нужно сравнивать только элементы.
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func main() {
	x := []int{1, 2, 3}
	y := []int{1, 2, 3}
	z := []int{1, 2, 4}

	fmt.Println(EqualInts(x, y))
	fmt.Println(EqualInts(x, z))
}
