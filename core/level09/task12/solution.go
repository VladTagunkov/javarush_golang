package main

import (
	"errors"
	"fmt"
	"strconv"
)

// parseExpr разбирает три токена "a op b" в типизированное выражение.
// Имена результатов в сигнатуре специально показывают смысл возвращаемых значений.
func parseExpr(aS, opS, bS string) (a int, op string, b int, err error) {
	aVal, err := strconv.Atoi(aS)
	if err != nil {
		return 0, "", 0, fmt.Errorf("cannot parse integer: %s", aS)
	}

	bVal, err := strconv.Atoi(bS)
	if err != nil {
		return 0, "", 0, fmt.Errorf("cannot parse integer: %s", bS)
	}
	switch opS {
	case "+":
		return aVal, opS, bVal, nil
	case "-":
		return aVal, opS, bVal, nil
	case "*":
		return aVal, opS, bVal, nil
	case "/":
		return aVal, opS, bVal, nil
	default:
		return aVal, opS, bVal, errors.New("unknown operation: " + opS)
	}

	// TODO: Проверьте, что opS — один из: +, -, *, /.
	// TODO: Если оператор неизвестен — верните ошибку (явным return со значениями).

	return aVal, opS, bVal, nil
}

// calc выполняет вычисление; неизвестный оператор и деление на ноль — ошибки.
func calc(a int, op string, b int) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		// TODO: Реализуйте вычитание.
		return a - b, nil
	case "*":
		// TODO: Реализуйте умножение.
		return a * b, nil
	case "/":
		// TODO: Реализуйте деление.
		// TODO: Верните ошибку при делении на ноль (b == 0).
		if b != 0 {
			return a / b, nil
		} else {
			return 0, errors.New("division by zero")
		}
	default:
		return 0, fmt.Errorf("unknown operator: %s", op)
	}
}

func main() {
	var aToken, opToken, bToken string
	fmt.Scan(&aToken, &opToken, &bToken)

	a, op, b, err := parseExpr(aToken, opToken, bToken)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	res, err := calc(a, op, b)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("result: %d\n", res)
}
