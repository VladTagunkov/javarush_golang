package main

import "fmt"

type Meta struct {
	ID int
}

type Task struct {
	Title string
	Meta  // TODO: Сделайте поле Meta embedded (анонимным), чтобы ID "поднялся" в Task.
}

func main() {
	task := Task{
		Title: "LearnEmbedding",
		Meta: Meta{
			ID: 42, // TODO: Инициализируйте ID = 42 через embedded Meta (promoted-поле), а не через явное поле Meta.
		},
	}

	// TODO: Выведите ID через короткий доступ task.ID (после embedding) — это должна быть первая строка.
	fmt.Println(task.ID)

	// Вторая строка должна выводить то же значение, но через явный путь task.Meta.ID.
	fmt.Println(task.Meta.ID)
}
