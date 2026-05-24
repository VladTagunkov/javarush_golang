package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	printClosure(n)
	printArgs(n)
}

func printClosure(n int) {

	fmt.Print("closure: ")
	defer fmt.Println()
	for i := 0; i < n; i++ {
		defer func() {
			fmt.Printf("%d", i)
		}()
	}
}

func printArgs(n int) {
	fmt.Print("args: ")
	defer fmt.Println()

	for i := 0; i < n; i++ {
		//_ = i
		// TODO: Зарегистрируйте defer для печати цифры через аргумент анонимной функции (как требуется в задаче).
		defer func(i_ int) {
			fmt.Printf("%d", i_)
		}(i)
	}
}
