package main

import "fmt"

func main() {
	var s, repl string
	fmt.Scan(&s, &repl)

	// Если s пустая — сразу печатаем пустую строку.
	if s == "" {
		fmt.Println("")
		return
	}

	// Если repl пустая — строку s не меняем.
	if repl == "" {
		fmt.Println(s)
		return
	}

	// TODO: Замените первый Unicode-символ строки s на первый Unicode-символ строки repl через преобразование в []rune.
	// TODO: Соберите результат строго через string(runes) и выведите ровно одну строку.
	rp := []rune(repl)
	sp := []rune(s)
	sp[0] = rp[0]
	s = string(sp)

	fmt.Println(s)
}
