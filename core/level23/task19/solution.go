package main

import "fmt"

type Note struct {
	Text string
}

type NoteList []Note

func (nl *NoteList) Add(text string) {
	*nl = append(*nl, Note{Text: text})
}

func (nl NoteList) Count() int {
	return len(nl)
}

type SafeNotebook struct {
	Title string
	notes NoteList
}

func (nb *SafeNotebook) AddNote(text string) {
	nb.notes.Add(text)
}

func (nb SafeNotebook) Notes() []Note {
	// TODO: Верните копию заметок (нужен новый слайс нужной длины и копирование элементов),
	// TODO: чтобы изменения во внешнем слайсе не влияли на внутреннее состояние блокнота.
	res := make([]Note, nb.Count())
	copy(res, nb.notes)
	return res
}

func (nb SafeNotebook) Count() int {
	return nb.notes.Count()
}

func main() {
	nb := SafeNotebook{Title: "My notes"}

	nb.AddNote("Buy")
	nb.AddNote("Milk")

	ext := nb.Notes()
	ext[0].Text = "HACK"

	internalFirst := nb.Notes()[0].Text

	fmt.Println("internal_first=" + internalFirst)
	fmt.Println("external_first=" + ext[0].Text)
	fmt.Printf("internal_count=%d\n", nb.Count())
}
