package main

import "fmt"

type Meta struct {
	ID      int
	Version int
}

type Document struct {
	Name string
	Meta // embedded как значение
}

func BumpVersion(d *Document) {
	// TODO: увеличьте версию документа на 1 через promoted-поле (короткий доступ)
	d.Version++
}

func main() {
	var id, version int
	var name string
	fmt.Scan(&id, &version, &name)

	doc := Document{
		// TODO: создайте документ через именованный литерал и заполните Meta данными из ввода
		Name: name,
		Meta: Meta{ID: id, Version: version},
	}

	BumpVersion(&doc)

	fmt.Printf("version=%d metaVersion=%d\n", doc.Version, doc.Meta.Version)
}
