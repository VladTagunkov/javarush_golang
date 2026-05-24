package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var (
		count int
		sum   int
	)

	for { // по требованию: бесконечный цикл с управлением через break
		var x int
		_, err := fmt.Fscan(in, &x)
		if err != nil { // конец ввода/ошибка чтения — просто выходим
			break
		}

		// команда остановки: в статистику не входит
		if x == 0 {
			break
		}

		if (x < -1000) || (x > 1000) {
			continue
		}

		if x == 777 {
			sum += x
			count++
			break
		} else {
			sum += x
			count++
		}

		// TODO: реализовать фильтрацию шумовых значений по условию задачи (с пропуском через continue)

		// TODO: реализовать учёт принятых значений (счётчик и сумма)

		// TODO: реализовать аварийную остановку по специальному сигналу из условия задачи
	}

	fmt.Println(count, sum)
}
