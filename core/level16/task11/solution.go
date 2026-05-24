package main

import (
	"errors"
	"fmt"
	"strconv"
)

// ParseNonNegativeInt парсит строку в int и запрещает отрицательные значения.
func ParseNonNegativeInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil { // важно: проверка сразу после Atoi
		// TODO: Вернуть ошибку парсинга как есть (без изменения текста) и 0 как значение.
		return 0, err
	}

	// TODO: Проверить, что число неотрицательное. Если n < 0 — вернуть 0 и отдельную (не парсинговую) ошибку.
	// Важно: текст ошибки должен быть строчным и без точки в конце.
	if n < 0 {
		return 0, errors.New("negative number")
	}
	return n, nil
}

func main() {
	var aStr, bStr string
	fmt.Scan(&aStr, &bStr)

	a, err := ParseNonNegativeInt(aStr)
	if err != nil {
		// TODO: Добавить контекст ошибки для aStr через fmt.Errorf с форматом: parse a %q: %v (без %w).
		//fmt.Printf("error: %v", err)
		err = fmt.Errorf("parse a %q: %v", aStr, err)
		fmt.Printf("error:%v", err)
		return
	}

	b, err := ParseNonNegativeInt(bStr)
	if err != nil {
		// TODO: Добавить контекст ошибки для bStr через fmt.Errorf с форматом: parse b %q: %v (без %w).
		//fmt.Printf("error: %v", err)
		err = fmt.Errorf("parse b %q: %v", bStr, err)
		fmt.Printf("error:%v", err)
		return
	}

	fmt.Printf("result: %d", a+b)
}
