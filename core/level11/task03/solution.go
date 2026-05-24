package main

import "fmt"

// Normalize принимает МАССИВ (копию), заменяет отрицательные на 0 и возвращает новый массив.
func Normalize(a [5]int) [5]int {
	// TODO: Реализуйте нормализацию массива: все элементы меньше 0 замените на 0.
	for i := range a {
		// TODO: Добавьте проверку значения a[i] и при необходимости замените его.
		if a[i] < 0 {
			a[i] = 0
		}
		//a[i] = a[i]
	}
	return a
}

func main() {
	var in [5]int

	// Читаем ровно 5 чисел в массив через fmt.Scan в цикле по индексам 0..4.
	for i := 0; i < 5; i++ {
		fmt.Scan(&in[i])
	}

	// Важно: массив не меняется "сам", результат нужно явно взять из return.
	out := Normalize(in)

	fmt.Printf("in: %v\n", in)
	fmt.Printf("out: %v\n", out)
}
