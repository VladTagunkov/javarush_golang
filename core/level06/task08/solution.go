package main

import "fmt"

func main() {
	var title string
	var planned int
	var spent float64
	var done bool
	var isBig bool
	var isLong bool
	var status string

	// Ввод строго в указанном порядке, title читаем до пробела.
	fmt.Scan(&title, &planned, &spent, &done)

	// TODO: Вычислите isBig по условию задачи на основе planned (сравнение int).
	if planned >= 5 {
		isBig = true
	} else {
		isBig = false
	}

	// TODO: Вычислите isLong по условию задачи на основе spent (сравнение float64).
	// Важно: используйте isLong только при расчёте needsAttention.
	if spent >= 3.0 {
		isLong = true
	} else {
		isLong = false
	}

	// TODO: Вычислите needsAttention по условию задачи, используя !, &&, || и скобки.
	// Подсказка по форме: needsAttention должен зависеть от done, isBig и isLong.
	needsAttention := (!done) && (isBig || isLong)

	// TODO: Сформируйте строку статуса через конкатенацию (+) в нужном формате
	// и в зависимости от done.

	if done {
		// TODO: Исправьте статус для случая done == true.
		status = title + ": DONE"
	} else {
		status = title + ": IN PROGRESS"
	}

	// Ровно три строки и в точном порядке.
	fmt.Println(status)
	fmt.Println("isBig=" + fmt.Sprint(isBig))
	fmt.Println("needsAttention=" + fmt.Sprint(needsAttention))
}
