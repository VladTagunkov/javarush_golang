package main

import (
	"errors"
	"fmt"
)

func SafeDiv(a int, b int) (int, error) {
	// TODO: Реализуйте безопасное целочисленное деление a / b.
	// TODO: Если делитель равен нулю — верните (0, error), где error создан через errors.New или fmt.Errorf.
	// TODO: Если ошибки нет — верните результат деления и nil.
	//_ = errors.New
	//return 0, nil
	if b == 0 {
		return 0, errors.New("Zero division.")
	}
	return a / b, nil
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	q, err := SafeDiv(a, b)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	fmt.Print(q)
}
