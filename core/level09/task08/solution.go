package main

import (
	"errors"
	"fmt"
	"strconv"
)

func parseIntStrict(s string) (int, error) {
	// TODO: Реализуйте строгий парсинг строки в int через strconv.
	// TODO: При любой ошибке парсинга возвращайте (0, non-nil error) с согласованным текстом ошибки.

	v, err := strconv.Atoi(s)
	if err != nil {
		// В шаблоне намеренно используется другой текст ошибки, чтобы решение не было готовым.
		return 0, errors.New("invalid integer: " + s)
	}

	// В шаблоне намеренно добавлено лишнее ограничение (запрет отрицательных),
	// чтобы поведение не соответствовало требованиям полностью.
	//if v < 0 {
	//	return 0, errors.New("invalid integer: negative values are not supported")
	//}

	return v, nil
}

func safeDivInt(a, b int) (int, error) {
	// TODO: Реализуйте безопасное целочисленное деление a/b.
	// TODO: При b == 0 возвращайте (0, errors.New(...)) с согласованным текстом ошибки.

	if b == 0 {
		// В шаблоне намеренно используется другой текст ошибки, чтобы решение не было готовым.
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func calcOneOp(a int, op string, b int) (int, error) {
	// TODO: Реализуйте вычисление одной операции по оператору op.
	// TODO: Поддержите +, -, *, /.
	// TODO: Для деления используйте safeDivInt (не делите напрямую).
	// TODO: Для неизвестного оператора возвращайте (0, non-nil error) с согласованным текстом ошибки.

	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":

		res, err := safeDivInt(a, b)
		return res, err
	default:
		// В шаблоне умышленно не реализованы * и /.
		return 0, errors.New("unknown operator")
	}
}

func main() {
	var aStr, op, bStr string
	fmt.Scan(&aStr, &op, &bStr)

	a, err := parseIntStrict(aStr)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	b, err := parseIntStrict(bStr)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	res, err := calcOneOp(a, op, b)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("result = %d\n", res)
}
