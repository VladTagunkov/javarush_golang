package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrBadInput = errors.New("bad input") // sentinel-ошибка

func parseCommand(cmd string) error {
	if cmd == "add" {
		return nil
	}

	// TODO: Верните ошибку "неизвестная команда", обернув ErrBadInput через fmt.Errorf с %w (с нужным текстом ошибки).
	return fmt.Errorf("unknown command %q: %w", cmd, ErrBadInput)
}

func run() error {
	var cmd string
	fmt.Fscan(os.Stdin, &cmd)

	if err := parseCommand(cmd); err != nil {
		// TODO: Добавьте верхний контекст для ошибки из parseCommand, обернув её через fmt.Errorf с %w (с нужным текстом ошибки).
		err2 := fmt.Errorf("run: %w", err)
		return err2
	}
	return nil
}

func printChain(err error) {
	if err == nil {
		return
	}
	for err != nil {
		fmt.Printf("- %v\n", err)
		err = errors.Unwrap(err)
	}

	// TODO: Допечатайте всю цепочку ошибок построчно, переходя к следующей причине через errors.Unwrap(err) до nil.
}

func main() {
	if err := run(); err != nil {
		printChain(err)
		return
	}
	fmt.Println("OK")
}
