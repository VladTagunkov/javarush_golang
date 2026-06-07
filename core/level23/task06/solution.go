package main

import "fmt"

type Meta struct {
	Title   string
	Author  string
	Project string
}

type Task struct {
	Meta
	// TODO: Сделайте Meta embedded-полем (анонимным), чтобы поля Meta "поднимались" в Task.
}

type StoredTask struct {
	ID int

	Task
	// TODO: Сделайте Task embedded-полем (анонимным), чтобы поля Task (и транзитивно Meta) "поднимались" в StoredTask.
}

func main() {
	var (
		id      int
		title   string
		author  string
		project string
	)

	fmt.Scan(&id, &title, &author, &project)

	stored := StoredTask{
		ID: id,
		Task: Task{
			Meta: Meta{
				Title:   title,
				Author:  author,
				Project: project,
			},
		},
	}

	// TODO: Выведите значения title/author/project через короткий доступ: stored.Title, stored.Author, stored.Project.
	fmt.Printf("id=%d; title=%s; author=%s; project=%s", stored.ID, stored.Title, stored.Author, stored.Project)
}
