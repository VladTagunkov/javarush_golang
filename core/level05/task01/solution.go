package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// TODO: Выведите ровно N строк в stdout.
	// TODO: Каждая строка должна иметь строгий формат: "Iteration <номер>" + перевод строки.
	// TODO: Нумерация должна начинаться с 1 и увеличиваться на 1 на каждой итерации.
	// TODO: Используйте классический цикл for с init/cond/post (счётчик объявите в init).
	// TODO: При N = 0 не выводите вообще ничего.
	if n == 0 {
		return
	}
	for i := 1; i <= n; i++ {
		fmt.Printf("Iteration %d\n", i)
	}
}
