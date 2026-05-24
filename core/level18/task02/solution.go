package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	demo(a, b)
}

func demo(a, b int) {
	draftTotal := a

	// TODO: Исправьте defer так, чтобы он печатал "saved=<число>" со значением,
	// сохранённым на момент постановки defer (а не тем, которое окажется в draftTotal позже).
	defer func(saved_ int) {
		fmt.Printf("saved=%d\n", saved_)
	}(a)

	draftTotal = b
	fmt.Printf("now=%d\n", draftTotal)
}
