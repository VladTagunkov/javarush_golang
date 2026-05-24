package main

import "fmt"

func main() {
	var currentValue int
	fmt.Scan(&currentValue)

	doublingsCount := 0

	// TODO: Реализуйте цикл for в форме "как while" (без init/post).
	// TODO: В цикле удваивайте currentValue, пока он строго меньше 1000, и увеличивайте doublingsCount на 1 на каждом шаге.
	// TODO: Обеспечьте гарантированное завершение программы для любых целых входных данных (в том числе currentValue <= 0).
	//if currentValue <= 0 {
	//	fmt.Println("0 0")
	//}
	for currentValue < 1000 {
		if currentValue <= 0 {
			currentValue = 0
			doublingsCount = 0
			break
		}
		currentValue *= 2
		doublingsCount++
	}

	fmt.Printf("%d %d", currentValue, doublingsCount)
}
