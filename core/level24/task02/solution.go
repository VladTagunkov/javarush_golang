package main

import (
	"fmt"
)

// Valuer — "источник числа": нам важно только уметь запросить значение.
type Valuer interface {
	Value() int
}

// NumberBox упаковывает одно число и реализует Valuer.
type NumberBox struct {
	N int
}

func (b NumberBox) Value() int {
	// TODO: Верните значение поля N, чтобы NumberBox корректно реализовывал интерфейс Valuer.
	return b.N
}

// Sum работает через интерфейс, а не через конкретные структуры.
func Sum(x, y Valuer) int {
	// TODO: Реализуйте сумму через вызовы Value() у x и y.
	return x.Value() + y.Value()
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	x := NumberBox{N: a}
	y := NumberBox{N: b}

	fmt.Println(Sum(x, y))
}
