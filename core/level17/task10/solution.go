package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ParseIntStrict(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		// TODO: Добавьте контекст и оберните исходную ошибку через %w так, чтобы причина strconv.Atoi сохранялась в цепочке.
		return 0, fmt.Errorf("parse int strict failed: %w", err)
	}

	// TODO: Верните распарсенное число и nil-ошибку при успешном парсинге.
	return n, nil
}

func main() {
	var s string
	fmt.Fscan(os.Stdin, &s)

	n, err := ParseIntStrict(s)
	if err == nil {
		fmt.Printf("NUMBER: %d\n", n)
		return
	}

	var numErr *strconv.NumError
	//_ = numErr
	if errors.As(err, &numErr) {
		fmt.Printf("NOT_A_NUMBER")
	} else {
		fmt.Printf("ERROR: %v\n", err)
	}

	// TODO: Реализуйте различение случая “не число” через errors.As(err, &numErr) и выведите NOT_A_NUMBER.
	// TODO: Во всех остальных случаях выведите ERROR: <текст ошибки>.

}
