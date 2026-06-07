package main

import (
	"errors"
	"fmt"
	"os"
)

type TitleValidator struct{}

func (TitleValidator) Validate(title string) error {
	if title == "" {
		return errors.New("title is empty")
	}

	// TODO: Добавьте проверку, что заголовок не длиннее 20 символов (len(title) > 20).
	// TODO: При нарушении верните ошибку с текстом: "title is too long".
	if len(title) > 20 {
		return errors.New("title is too long")
	}
	return nil
}

type TaskService struct {
	validator TitleValidator
}

func (s TaskService) NewTask(title, author, project string) (Task, error) {
	if err := s.validator.Validate(title); err != nil {
		return Task{}, err
	}

	// TODO: Создайте задачу так, чтобы в ней были корректно заполнены метаданные (author/project),
	// TODO: но при этом поля метаданных оставались инкапсулированными (неэкспортируемыми).
	return Task{
		Title: title,
		meta:  Meta{author: author, project: project},
	}, nil
}

type Meta struct {
	author  string // неэкспортируемые поля: снаружи нельзя менять напрямую
	project string
}

func (m Meta) Author() string {
	// TODO: Верните автора из метаданных.
	return m.author
}

func (m Meta) Project() string {
	// TODO: Верните проект из метаданных.
	return m.project
}

type Task struct {
	Title string
	meta  Meta // метаданные спрятаны внутри задачи
}

func (t Task) Author() string {
	// TODO: Реализуйте метод-обёртку: верните автора через t.meta (без прямого доступа к полям meta снаружи).
	return t.meta.Author()
}

func (t Task) Project() string {
	// TODO: Реализуйте метод-обёртку: верните проект через t.meta (без прямого доступа к полям meta снаружи).
	return t.meta.Project()
}

func main() {
	var title, author, project string
	fmt.Fscan(os.Stdin, &title, &author, &project)

	svc := TaskService{validator: TitleValidator{}}

	task, err := svc.NewTask(title, author, project)
	if err != nil {
		fmt.Printf("validation: %s\n", err)
		return
	}

	fmt.Printf("created: %s; author=%s; project=%s\n", task.Title, task.Author(), task.Project())
}
