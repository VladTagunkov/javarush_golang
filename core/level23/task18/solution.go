package main

import "fmt"

type Note struct {
	Text string
}

// Именованный тип поверх слайса заметок.
type NoteList []Note

func (nl *NoteList) Add(text string) {
	// TODO: Реализуйте добавление новой заметки в список (через append).
	// TODO: Сохраните текст заметки в поле Note.Text.
	*nl = append(*nl, Note{text})
}

func (nl NoteList) Count() int {
	// TODO: Реализуйте подсчёт количества заметок в списке.
	return len(nl)
}

type Notebook struct {
	Title string
	// Embedding "протаскивает" наружу и методы, и само поле NoteList (его можно заменить целиком).
	NoteList
}

func main() {
	nb := Notebook{Title: "My notes"}
	nb.Add("First")
	nb.Add("Second")
	// TODO: Добавьте в блокнот две заметки через nb.Add(...).

	fmt.Printf("before=%d\n", nb.Count())

	// TODO: Сбросьте список заметок присваиванием nb.NoteList = NoteList{}.
	nb.NoteList = NoteList{}
	fmt.Printf("after=%d\n", nb.Count())
}
