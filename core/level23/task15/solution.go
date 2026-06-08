package main

import "fmt"

type LeftTag struct{}

func (LeftTag) Tag() string {
	// TODO: Реализуйте метод Tag() для LeftTag по условию задачи.
	return "LEFT"
}

type RightTag struct{}

func (RightTag) Tag() string {
	// TODO: Реализуйте метод Tag() для RightTag по условию задачи.
	return "RIGHT"
}

type Pair struct {
	LeftTag
	RightTag
}

func main() {
	pair := Pair{}

	// TODO: Исправьте вывод так, чтобы вызывались оба метода Tag() явно через embedded-типы.
	// Вызов pair.Tag() здесь оставлен намеренно — его нужно заменить на корректные явные вызовы.
	fmt.Printf("%s %s\n", pair.LeftTag.Tag(), pair.RightTag.Tag())
}
