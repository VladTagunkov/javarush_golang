package main

import "fmt"

// apply — чистая функция: не читает и не печатает, только считает новое значение.
func apply(op string, current int, x int) int {
	// TODO: Реализуйте логику операций add/sub/set.
	// TODO: При неизвестной операции верните current без изменений.
	switch op {
	case "add":
		return current + x
	case "sub":
		return current - x
	case "set":
		return x
	default:
		return current
	}
	//return current
}

func main() {
	current := 0 // по требованиям — только локальное состояние, без глобальных переменных
	var op string
	var x int

	// TODO: Реализуйте цикл чтения пар (операция, число) через fmt.Scan.
	// TODO: Немедленно завершайте цикл при команде stop 0 (без вызова apply).
	// TODO: Для остальных команд обновляйте current через apply.
	for i := 0; true; i++ {
		if _, err := fmt.Scanf("%s %d", &op, &x); err != nil {
			break
		}
		if op == "stop" && x == 0 {
			break
		}
		//fmt.Scanf(&op, &current, &x)
		current = apply(op, current, x)
	}

	// Ровно одна строка с итогом.
	fmt.Println(current)
}
