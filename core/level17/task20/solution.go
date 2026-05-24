package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Sentinel-ошибки: классы проблем, распознаём только через errors.Is.
var (
	ErrValidation = errors.New("validation")
	ErrNotFound   = errors.New("notfound")
	ErrIO         = errors.New("io")
	ErrInternal   = errors.New("internal")
)

func main() {
	// Заранее заданный список задач.
	tasks := []string{
		"купить молоко",
		"прочитать про errors.Is",
		"сделать домашку",
	}

	var cmd, arg string
	if _, err := fmt.Scan(&cmd, &arg); err != nil {
		// TODO: Классифицируйте ошибку чтения как io через errors.Join(ErrIO, err) и оберните контекстом так,
		// TODO: чтобы errors.Is(err, ErrIO) срабатывал даже при дополнительном wrapping.
		printUser(fmt.Errorf("read cmd/arg: %w", ErrIO))
		return
	}

	res, err := dispatch(cmd, arg, tasks)
	if err != nil {
		printUser(err)
		return
	}

	fmt.Printf("RESULT %s\n", res)
}

func dispatch(cmd, arg string, tasks []string) (string, error) {
	switch cmd {
	case "get":
		// TODO: Реализуйте команду "get":
		// TODO: - распарсить id из arg (strconv.Atoi) и отнести ошибки к validation (можно через errors.Join)
		// TODO: - id < 0 -> validation
		// TODO: - id вне диапазона tasks -> notfound (через wrapping с %w)
		// TODO: - при успехе вернуть tasks[id]
		return "", fmt.Errorf("get: %w", ErrInternal)

	case "first":
		// TODO: Реализуйте команду "first":
		// TODO: - при пустом tasks вернуть ошибку класса internal
		// TODO: - иначе вернуть первую задачу
		return "", fmt.Errorf("first: %w", ErrInternal)

	case "count":
		return strconv.Itoa(len(tasks)), nil

	default:
		return "", fmt.Errorf("unknown command %q: %w", cmd, ErrValidation)
	}
}

func printUser(err error) {
	// Сообщение выбираем строго по errors.Is (никаких разборов текста ошибки).
	msg := "Внутренняя ошибка"

	switch {
	case errors.Is(err, ErrValidation):
		msg = "Неверный ввод"
		// TODO: Добавьте остальные классы ошибок (notfound/io/internal) и их сообщения, используя только errors.Is.
	}

	fmt.Printf("USER %s\n", msg)
}