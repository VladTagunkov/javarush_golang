package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	// TODO: Вычислите sum, diff и prod для целых чисел a и b с помощью операторов +, -, *.
	sum := a + b
	diff := a - b
	prod := a * b

	fmt.Printf("sum=%d\n", sum)
	fmt.Printf("diff=%d\n", diff)
	fmt.Printf("prod=%d\n", prod)

	// TODO: Выведите div и mod в строгом формате.
	// TODO: Если b == 0, нельзя выполнять a/b и a%b — нужно вывести div=undefined и mod=undefined.
	// TODO: Если b != 0, нужно вывести целочисленные результаты a/b и a%b.
	if b == 0 {
		fmt.Print("div=undefined\n")
		fmt.Print("mod=undefined\n")
	} else {
		diff := a / b
		mod := a % b
		fmt.Printf("div=%d\n", diff)
		fmt.Printf("mod=%d\n", mod)
	}

}
