package main

import (
	"errors"
	"fmt"
)

type Status string

const (
	StatusTodo Status = "todo"
	StatusDone Status = "done"
)

// Логика "задача сделана?" живёт рядом со статусом.
func (s Status) IsDone() bool {
	// TODO: Реализуйте логику определения "сделано ли" для статуса.
	if s == StatusDone {
		return true
	}
	return false
}

type Task struct {
	Title  string
	Status Status
}

// Task не знает правил done — просто делегирует статусу.
func (t Task) Done() bool {
	return t.Status.IsDone()
}

var errInvalidStatus = errors.New("invalid status")

func ParseStatus(text string) (Status, error) {
	// TODO: Реализуйте разбор текста статуса: принимайте только "todo" и "done", иначе возвращайте errInvalidStatus.

	// Частичная реализация для каркаса: распознаём только "todo".
	if text == string(StatusTodo) {
		return StatusTodo, nil
	} else if text == string(StatusDone) {
		return StatusDone, nil
	}

	return "", errInvalidStatus
}

func main() {
	var title, statusText string
	fmt.Scan(&title, &statusText)

	status, err := ParseStatus(statusText)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	task := Task{
		Title:  title,
		Status: status,
	}

	fmt.Printf("done=%t; title=%s", task.Done(), task.Title)
}
