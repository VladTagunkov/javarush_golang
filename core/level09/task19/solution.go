package main

import "fmt"

// makeStepper возвращает функцию-счётчик.
//
// ВАЖНО: состояние счётчика должно храниться внутри замыкания.
// current := 0
func makeStepper(start, step int) func() int {
	// TODO: Реализуйте хранение текущего значения внутри замыкания.
	// TODO: При каждом вызове возвращайте очередное значение, начиная со start,
	// TODO: и изменяйте внутреннее состояние на step для следующего вызова.

	current := start
	//v := current
	return func() int {
		// TODO: Верните текущее значение счётчика и обновите его для следующего тика.
		toReturn := current
		current += step
		return toReturn
	}
}

func main() {
	var startValue, stepSize, ticksCount int
	fmt.Scan(&startValue, &stepSize, &ticksCount)

	if ticksCount <= 0 {
		return
	}

	next := makeStepper(startValue, stepSize)
	for i := 0; i < ticksCount; i++ {
		fmt.Println(next())
	}
}
