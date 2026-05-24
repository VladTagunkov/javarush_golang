package main

import "fmt"

func main() {
	var x int
	var s string

	fmt.Scan(&x, &s)

	// TODO: Выведите значение x и его тип в требуемом формате через fmt.Printf (без fmt.Println и без лишнего текста).
	fmt.Printf("x=%v type=%T\n", x, x)

	// TODO: Выведите значение s и его тип в требуемом формате через fmt.Printf (без fmt.Println и без лишнего текста).
	fmt.Printf("s=%v type=%T\n", s, s)
}
