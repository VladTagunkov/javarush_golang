package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	// TODO: Сделайте так, чтобы b и a указывали на один и тот же набор элементов (без копирования).
	// Сейчас здесь создаётся отдельный слайс, и изменения через b не отражаются в a.
	//b = append([]int(nil), a...)

	b[0] = 99

	fmt.Printf("a=%v\n", a)
	fmt.Printf("b=%v\n", b)

	// TODO: Выведите len и cap для a и b в точном формате из задания.
	fmt.Printf("a: len=%d cap=%d; b: len=%d cap=%d\n", len(a), cap(a), len(b), cap(b))
}
