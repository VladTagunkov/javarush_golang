package main

import "fmt"

type Ident struct {
	ID int
}

type Audit struct {
	ID        int
	CreatedBy string
}

type Ticket struct {
	Title string
	Ident // embedded
	Audit // embedded
}

func (t Ticket) Summary() string {
	// TODO: Реализуйте формирование строки сводки по заданному формату.
	// TODO: Из-за конфликта имён поля ID в embedded-структурах обращайтесь к полям только явными путями: t.Ident.ID и t.Audit.ID.
	// TODO: Строка должна строго соответствовать шаблону: title=<Title> ident=<Ident.ID> audit=<Audit.ID> by=<Audit.CreatedBy>
	return fmt.Sprintf("title=%s ident=%d audit=%d by=%s", t.Title, t.Ident.ID, t.Audit.ID, t.Audit.CreatedBy)
}

func main() {
	var title string
	var identID int
	var auditID int
	var createdBy string

	fmt.Scan(&title, &identID, &auditID, &createdBy)

	var ticket Ticket
	// TODO: Корректно заполните все поля ticket считанными значениями (Title, Ident.ID, Audit.ID, Audit.CreatedBy), используя явные пути.
	ticket.Title = title
	ticket.Ident = Ident{ID: identID}
	ticket.Audit = Audit{ID: auditID, CreatedBy: createdBy}
	fmt.Print(ticket.Summary())
}
