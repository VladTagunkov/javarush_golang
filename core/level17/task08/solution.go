package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Sentinel-ошибки = "классы" ошибок, которые распознаются через errors.Is.
var (
	ErrUnknownOp = errors.New("unknown op")
	ErrDivByZero = errors.New("division by zero")
)

func calc(op, aStr, bStr string) (int, error) {
	// TODO: Реализуйте проверку операции так, чтобы неизвестная операция возвращала ErrUnknownOp
	// TODO: Важно: неизвестную операцию нужно обрабатывать сразу, без парсинга чисел

	// Упрощённая (неполная) проверка операции.
	if op != "add" && op != "sub" && op != "mul" && op != "div" {
		return 0, ErrUnknownOp
	}

	a, err := strconv.Atoi(aStr)
	if err != nil {
		// TODO: Оберните ошибку парсинга через fmt.Errorf(... %w ...), чтобы *strconv.NumError оставался доступен через errors.As
		return 0, fmt.Errorf("parse %q: %w", aStr, err)
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		// TODO: Оберните ошибку парсинга через fmt.Errorf(... %w ...)
		return 0, fmt.Errorf("parse %q: %w", bStr, err)
	}

	// TODO: Реализуйте вычисления для add/sub/mul/div
	// TODO: Для div при b == 0 верните ErrDivByZero (класс ошибки через sentinel)
	// Сейчас намеренно неполная логика: возвращаем сложение для любой операции.
	if op == "add" {
		return a + b, nil
	} else if op == "sub" {
		return a - b, nil
	} else if op == "mul" {
		return a * b, nil
	} else if op == "div" {
		if b == 0 {
			return 0, ErrDivByZero
		}
		return a / b, nil
	}
	return 0, nil
}

func main() {
	var op, aStr, bStr string
	fmt.Fscan(os.Stdin, &op, &aStr, &bStr)

	res, err := calc(op, aStr, bStr)
	if err == nil {
		fmt.Println(res)
		return
	}

	// TODO: На границе приложения нужно оборачивать ошибку через fmt.Errorf("calc: %w", err)
	// Сейчас обёртка намеренно сделана неверно (без %w), чтобы терялась цепочка причин.
	//err = fmt.Errorf("calc: %v", err)
	err = fmt.Errorf("calc: %w", err)

	// Частично реализовано: класс неизвестной операции через errors.Is.
	// TODO: Дополните обработку ErrDivByZero через errors.Is
	// TODO: Дополните обработку BAD_NUMBER через errors.As и *strconv.NumError
	if errors.Is(err, ErrUnknownOp) {
		fmt.Println("UNKNOWN_OP")
		return
	} else if errors.Is(err, ErrDivByZero) {
		fmt.Println("DIV_BY_ZERO")
		return
	}
	var numErr *strconv.NumError
	if errors.As(err, &numErr) {
		fmt.Printf("BAD_NUMBER: %s", numErr.Num)
		return
	}
	fmt.Printf("ERROR: %v\n", err)
}
