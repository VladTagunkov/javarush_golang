package main

import "fmt"

// IDGen хранит состояние генератора ID.
type IDGen struct {
	next int
}

// Next возвращает следующий ID и обновляет состояние генератора.
func (g *IDGen) Next() int {
	// TODO: Реализуйте генерацию следующего ID: увеличить g.next на 1 и вернуть новое значение (метод должен менять состояние).
	g.next = g.next + 1
	return g.next
}

// Project встраивает IDGen, чтобы Next() был promoted-методом.
type Project struct {
	Name string
	IDGen
}

func main() {
	p := Project{Name: "Home"}

	fmt.Println(p.Name)
	fmt.Println(p.Next())
	fmt.Println(p.Next())
	fmt.Println(p.Next())
}
