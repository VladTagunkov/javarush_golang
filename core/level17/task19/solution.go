package main

import (
	"errors"
	"fmt"
)

// Sentinel-ошибки классов (создаются один раз на уровне пакета).
var (
	ErrValidation = errors.New("validation")
	ErrNotFound   = errors.New("not found")
	ErrInternal   = errors.New("internal")
)

func LookupColor(palette map[string]string, name string) (string, error) {
	// TODO: Реализуйте поиск кода цвета по имени.
	// TODO: Добавьте проверки palette == nil, name == "" и отсутствие ключа в palette.
	// TODO: Для отсутствия ключа используйте ok-идиому: v, ok := palette[name].
	// TODO: Возвращайте ошибки классов через fmt.Errorf с контекстом и %w, чтобы работал errors.Is.
	if palette == nil {
		return "", fmt.Errorf("lookup color: %w", ErrInternal)
	}
	if name == "" {
		return "", fmt.Errorf("string name: %w", ErrValidation)
	}
	v, ok := palette[name]
	if !ok {
		return "", fmt.Errorf("find palett: %w", ErrNotFound)
	}

	return v, nil

}

func main() {
	palette := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	var name string
	fmt.Scan(&name)

	code, err := LookupColor(palette, name)
	if err == nil {
		fmt.Println("COLOR", code)
		return
	}

	// TODO: Определите класс ошибки через errors.Is (validation/notfound/internal)
	// TODO: и напечатайте "ERR " + класс одним словом (без panic).
	if errors.Is(err, ErrInternal) {
		fmt.Println("ERR", "internal")
	}
	if errors.Is(err, ErrNotFound) {
		fmt.Println("ERR", "not found")
	}
	if errors.Is(err, ErrValidation) {
		fmt.Println("ERR", "validation")
	}
	//fmt.Println("ERR", "internal")
}
