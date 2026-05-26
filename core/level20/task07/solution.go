package main

import (
	"fmt"
	"strings"
)

func main() {
	var line string
	fmt.Scan(&line) // читаем одну "склеенную" строку (один токен)

	line = strings.ReplaceAll(line, "_", " ")
	normalized := strings.TrimSpace(line)

	if normalized == "" {
		fmt.Println("error: empty command")
		return
	}

	// TODO: Нормализуйте пробелы строго через связку strings.TrimSpace + strings.Fields (получите parts).
	// TODO: Если после нормализации parts пуст, выведите ровно одну строку: error: empty command и завершитесь.
	// TODO: Команда — это первое слово, приведите её к нижнему регистру через strings.ToLower.
	// TODO: Аргументы — это остальные слова; выведите их через запятую (strings.Join(args, ",")) в нужном формате.
	normalized = strings.TrimSpace(normalized)
	parts := strings.Fields(normalized)
	if len(parts) == 0 {
		fmt.Println("error: empty command")
		return
	}
	cmd := strings.ToLower(parts[0])
	args := parts[1:]
	// Временная упрощённая логика: считаем всю строку командой и игнорируем аргументы.
	// Это не соответствует требованиям задачи и должно быть исправлено.
	//cmd := strings.ToLower(normalized)
	//args := []string{}

	fmt.Println("cmd=" + cmd)
	fmt.Printf("argc=%d\n", len(args))
	fmt.Println("args=" + strings.Join(args, ","))
}
