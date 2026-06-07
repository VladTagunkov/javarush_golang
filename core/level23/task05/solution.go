package main

import "fmt"

type Meta struct {
	// TODO: Добавьте поля Author и Project типа string.
	Author  string
	Project string
}

type Task struct {
	// TODO: Добавьте поле Title типа string.
	// TODO: Добавьте именованное поле Meta типа Meta (не embedded), чтобы доступ был только через task.Meta.<Field>.
	Title string
	Meta  Meta
}

func main() {
	// TODO: Прочитайте из stdin три слова (title, author, project) с помощью fmt.Scan.
	// TODO: Создайте значение task типа Task, записав:
	// TODO: - title в task.Title
	// TODO: - author в task.Meta.Author
	// TODO: - project в task.Meta.Project
	// TODO: Выведите ровно одну строку в формате:
	// TODO: title=<TITLE>; author=<AUTHOR>; project=<PROJECT>
	// TODO: При выводе обращайтесь к метаданным явно: task.Meta.Author и task.Meta.Project.
	var title string
	var author string
	var project string
	fmt.Scan(&title, &author, &project)
	task := Task{}
	task.Title = title
	task.Meta.Author = author
	task.Meta.Project = project

	fmt.Printf("title=%s; author=%s; project=%s\n", task.Title, task.Meta.Author, task.Meta.Project)
}
