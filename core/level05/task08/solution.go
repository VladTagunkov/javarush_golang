package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for { // по условию: обработка в бесконечном цикле
		var cmd, a, b int
		if _, err := fmt.Fscan(in, &cmd, &a, &b); err != nil {
			break // EOF/ошибка ввода — корректно завершаемся
		}

		if cmd == 0 {
			break // команда завершения
		}
		if cmd == 1 {
			fmt.Fprintf(out, "%d\n", a+b)
		} else if cmd == 2 {
			fmt.Fprintf(out, "%d\n", a-b)
		} else if cmd == 3 {
			fmt.Fprintf(out, "%d\n", a*b)
		} else if cmd == 4 {
			if b != 0 {
				fmt.Fprintf(out, "%d\n", a/b)
			} else {
				continue
			}
		} else {
			continue
		}
		// TODO: реализуйте игнорирование неверных команд (cmd вне диапазона 0..4) через continue

		// TODO: реализуйте выполнение арифметических команд 1..4 и построчный вывод результата в stdout

		// TODO: реализуйте пропуск деления на ноль (cmd == 4 и b == 0) через continue без вывода

		// Временное поведение: корректно читаем и завершаемся по cmd==0/EOF,
		// но команды пока не обрабатываем.
		//continue
	}
}
