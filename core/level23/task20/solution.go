package main

import (
	"errors"
	"fmt"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

type IDGen struct {
	next int
}

func (g *IDGen) Next() int {
	// TODO: Реализуйте генерацию следующего ID (увеличить внутреннее состояние и вернуть новое значение).
	g.next = g.next + 1
	return g.next
}

type Project struct {
	Name  string
	tasks []Task // не embedded: наружу не должен торчать внутренний слайс
	idGen IDGen  // не embedded: наружу не должен торчать генератор
}

func (p *Project) AddTask(title string) (Task, error) {
	if title == "" {
		return Task{}, errors.New("empty title")
	}

	// TODO: Создайте задачу с корректным ID (через генератор), добавьте её во внутреннее хранилище и верните её наружу.
	t := Task{
		ID:    p.idGen.Next(),
		Title: title,
		Done:  false,
	}
	p.tasks = append(p.tasks, t)
	return t, nil
}

func (p *Project) SetDone(id int) bool {
	// TODO: Найдите задачу по ID во внутреннем хранилище, выставьте Done=true и верните true. Если не нашли — верните false.
	for item := range p.tasks {
		if p.tasks[item].ID == id {
			p.tasks[item].Done = true
			return true
		}
	}
	return false
}

func (p *Project) Tasks() []Task {
	// TODO: Верните копию слайса задач, чтобы внешний код не мог изменить внутреннее хранилище проекта.
	res := make([]Task, len(p.tasks))
	copy(res, p.tasks)
	return res
}

func main() {
	var p Project
	p.Name = "demo"

	t1, _ := p.AddTask("milk")
	fmt.Printf("task:%d %s\n", t1.ID, t1.Title)

	t2, _ := p.AddTask("room")
	fmt.Printf("task:%d %s\n", t2.ID, t2.Title)

	_, err := p.AddTask("")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}

	ok := p.SetDone(2)
	if ok {
		fmt.Println("done=1")
	} else {
		fmt.Println("done=0")
	}
}
