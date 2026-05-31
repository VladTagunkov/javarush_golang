package main

import "fmt"

type Message struct {
	// TODO: объявите поля структуры: From string, To string, Code int (имена и регистр важны)
	From, To string
	Code     int
}

func main() {
	var fromLogin, toLogin string
	var messageCode int

	fmt.Scan(&fromLogin, &toLogin, &messageCode)
	var m Message
	m.From = fromLogin
	m.To = toLogin
	m.Code = messageCode
	// TODO: создайте переменную структуры строго через `var m Message` (без литерала структуры)
	// TODO: заполните поля структуры отдельными присваиваниями: m.From = ..., m.To = ..., m.Code = ...
	// TODO: выведите ровно одну строку отчёта в требуемом формате (строки через %q, код как число)
	fmt.Printf("From=%q To=%q Code=%d\n", m.From, m.To, m.Code)
}
