package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Scan(&name, &age)

	// TODO: Выведите приветствие строго в формате из условия, используя fmt.Printf, %s (для имени) и %d (для возраста), и перевод строки в конце.
	// Важно: вывод должен быть ровно одной строкой, без лишних пробелов и дополнительных строк.
	fmt.Printf("Hello, %s! Age=%d\n", name, age)
}
