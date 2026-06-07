package main

import "fmt"

type Helloer interface {
	Hello(name string) string
}

type Byer interface {
	Bye(name string) string
}

// Greeter собирается через embedding маленьких интерфейсов, а не через перечисление методов.
type Greeter interface {
	Helloer
	Byer
}

type SimpleGreeter struct{}

func (SimpleGreeter) Hello(name string) string {
	// TODO: Реализуйте формирование строки приветствия с использованием имени.
	return fmt.Sprintf("Hello %s", name)
}

func (SimpleGreeter) Bye(name string) string {
	// TODO: Реализуйте формирование строки прощания с использованием имени.

	return fmt.Sprintf("Bye %s", name)
}

func main() {
	var userName string
	fmt.Scan(&userName)

	// Важно: используем переменную интерфейсного типа для вызова методов.
	var g Greeter = SimpleGreeter{}

	fmt.Println(g.Hello(userName))
	fmt.Println(g.Bye(userName))
}
