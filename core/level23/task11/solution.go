package main

import (
	"errors"
	"fmt"
	"os"
)

type Note struct {
	ID   int
	Text string
}

type NoteAdder interface {
	Add(n Note) error
}

type NoteFinder interface {
	FindByID(id int) (Note, bool)
}

type NoteStore interface {
	NoteAdder
	NoteFinder
}

type MemNoteStore struct {
	notes []Note
}

func (s *MemNoteStore) Add(n Note) error {
	// TODO: Реализуйте добавление заметки в in-memory хранилище.
	// TODO: Проверьте уникальность ID: если такой ID уже есть в хранилище — верните error (panic использовать нельзя).

	// Временная упрощённая проверка: сравнение только с последней заметкой.
	// Этого недостаточно, если дубликат ID встречается не подряд.
	if len(s.notes) > 0 {
		for _, nt := range s.notes {
			if nt.ID == n.ID {
				return errors.New("id already exists")
			}
		}

	}
	s.notes = append(s.notes, n)
	return nil
}

func (s *MemNoteStore) FindByID(id int) (Note, bool) {
	// TODO: Реализуйте поиск заметки по заданному id по всему хранилищу.
	// TODO: Верните (Note, true) при успешном нахождении, иначе (нулевую Note, false).

	if len(s.notes) == 0 {
		return Note{}, false
	}

	// Временная упрощённая реализация: проверяем только последнюю заметку.
	//last := s.notes[len(s.notes)-1]
	//if last.ID == id {
	//	return last, true
	//}
	for _, nt := range s.notes {
		if nt.ID == id {
			return nt, true
		}
	}
	return Note{}, false
}

func main() {
	in := os.Stdin

	var store NoteStore = &MemNoteStore{}

	for i := 0; i < 3; i++ {
		var id int
		var text string
		fmt.Fscan(in, &id, &text)

		if err := store.Add(Note{ID: id, Text: text}); err != nil {
			fmt.Printf("error: %s\n", err)
		}
	}

	var queryID int
	fmt.Fscan(in, &queryID)

	if n, ok := store.FindByID(queryID); ok {
		fmt.Printf("found: %d %s\n", n.ID, n.Text)
	} else {
		fmt.Println("not found")
	}
}
