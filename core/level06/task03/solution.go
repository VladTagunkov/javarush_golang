package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("0 0")
		return
	}

	// "Накопители" живут дольше цикла: объявляем снаружи.
	var sum int
	var max int

	for i := 0; i < n; i++ {
		// Текущее число — внутри тела цикла, чтобы показать область видимости.
		var x int
		fmt.Scan(&x)
		if i == 0 {
			max = x
		}

		sum = sum + x
		if x > max {
			max = x
		}

		// TODO: увеличьте сумму на x через обычное присваивание (оператор =) без использования слайсов/массивов
		// TODO: найдите максимум среди введённых чисел; корректно инициализируйте max без "магических" значений
	}

	fmt.Println(sum, max)
}
