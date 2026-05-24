package main

import "fmt"

func main() {
	sumParts := 0.0
	expected := 0.0

	sumParts = 0.1 + 0.2
	expected = 0.3
	// TODO: Вычислите sumParts как сумму 0.1 и 0.2 (тип float64).
	// TODO: Присвойте expected значение 0.3 (тип float64).

	fmt.Printf("%.17f\n", sumParts)
	fmt.Printf("%.17f\n", expected)

	eq := sumParts == expected
	// TODO: Выполните сравнение sumParts == expected оператором == и сохраните результат в eq.

	fmt.Println(eq)

	diff := sumParts - expected
	// TODO: Вычислите diff как (sumParts - expected).
	fmt.Printf("%.17f\n", diff)
}
