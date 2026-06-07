package main

import "fmt"

// Узкий контракт: нужно только имя.
type Namer interface {
	Name() string
}

// Узкий контракт: нужен только возраст.
type Ager interface {
	Age() int
}

// Сборка более широкого контракта через embedding интерфейсов.
type Profile interface {
	Namer
	Ager
}

type Student struct {
	name string
	age  int
}

func (s Student) Name() string {
	// TODO: Верните имя студента из структуры Student.
	return s.name
}

func (s Student) Age() int {
	// TODO: Верните возраст студента из структуры Student.
	return s.age
}

func PrintName(n Namer) {
	// TODO: Напечатайте только имя по контракту Namer (не используя Age) строго в формате:
	// name=<имя>
	fmt.Println("name=" + n.Name())
}

func PrintProfile(p Profile) {
	// TODO: Напечатайте имя и возраст по контракту Profile строго в формате:
	// name=<имя> age=<возраст>
	fmt.Println("name=" + p.Name() + " age=" + fmt.Sprintf("%d", p.Age()))
}

func main() {
	var studentName string
	var studentAge int
	fmt.Scan(&studentName, &studentAge)

	student := Student{name: studentName, age: studentAge}

	PrintName(student)
	PrintProfile(student)
}
