package main

import "fmt"

func main() {
	var a [3]int
	fmt.Scan(&a[0], &a[1], &a[2])

	//var b [3]int
	// TODO: Создайте массив b как копию массива a с помощью простого присваивания (без циклов и без copy).
	// TODO: Измените только массив b: увеличьте b[0] в 10 раз так, чтобы массив a не изменился.
	b := a
	b[0] = b[0] * 10
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
