package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {
	var id int
	var name string

	// Читаем входные данные (id и name без пробелов).
	_, _ = fmt.Scan(&id, &name)

	u1 := User{ID: id, Name: name}
	u2 := User{ID: id, Name: name}

	//u1 := &p1
	//u2 := &p2

	ptrEqual := &u1 == &u2
	valueEqual := u1 == u2
	// TODO: Создайте две независимые переменные типа User с одинаковыми данными из ввода (ID и Name).
	// TODO: Создайте два указателя на эти разные переменные (p1 и p2).
	// TODO: Выполните сравнение указателей (p1 == p2) и сравнение значений (*p1 == *p2).
	// TODO: Выведите ровно две строки в формате:
	// TODO: PTR_EQUAL=<значение>
	// TODO: VALUE_EQUAL=<значение>

	// Временная реализация: сохраняет формат вывода, но не реализует требуемую логику сравнения.
	//ptrEqual := false
	//valueEqual := false

	fmt.Printf("PTR_EQUAL=%v\n", ptrEqual)
	fmt.Printf("VALUE_EQUAL=%v\n", valueEqual)
}
