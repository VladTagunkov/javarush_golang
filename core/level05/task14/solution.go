package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("empty")
		return
	}
	for i := 1; i <= n; i++ {
		fmt.Printf("step %d\n", i)
	}
	// TODO: Реализуйте "тренажёр шагов": выведите ровно n строк в формате `step K`,
	// TODO: где K начинается с 1 и увеличивается на 1 до n включительно.
	// TODO: Используйте классический цикл `for init; cond; post` на ровно n итераций.
}
