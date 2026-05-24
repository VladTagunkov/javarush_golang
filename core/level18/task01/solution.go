package main

import "fmt"

func checkAge(age int) string {
	fmt.Println("start")
	defer fmt.Println("finish")
	// TODO: Отложите печать строки "finish" через defer так, чтобы она печаталась при любом выходе из функции.
	if age < 0 {
		return "invalid"
	}
	// TODO: Реализуйте проверку возраста:
	// TODO: - если age < 0, верните "invalid" (ранним return)
	// TODO: - иначе верните "ok"
	return "ok"
}

func main() {
	var age int
	fmt.Scan(&age)

	res := checkAge(age)
	fmt.Printf("result: %s\n", res)
}
