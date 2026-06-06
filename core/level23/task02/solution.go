package main

import "fmt"

// Meta — "карточка" с ID.
type Meta struct {
	ID int
}

// Note — заметка с текстом и метаданными.
type Note struct {
	Text string
	Meta // TODO: Сделайте Meta embedded-полем (анонимным), чтобы ID был доступен как note.ID
}

func main() {
	var noteID int
	var noteText string

	fmt.Scan(&noteID, &noteText)

	// TODO: Создайте переменную note типа Note через именованный литерал
	// TODO: Явно заполните Text и метаданные (ID должен попасть внутрь Meta)
	note := Note{Text: noteText, Meta: Meta{ID: noteID}}

	// TODO: Выведите ровно одну строку строго в формате: Note <id>: <text>
	// TODO: ID должен браться коротким доступом note.ID (а не note.Meta.ID)
	fmt.Printf("Note %d: %s", note.ID, note.Text)
}
