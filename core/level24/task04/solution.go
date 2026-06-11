package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// PositiveIntError — typed error: хранит исходный ввод и причину.
type PositiveIntError struct {
	Input   string
	Problem string
}

// Error реализует интерфейс error.
// TODO: Сформируйте сообщение об ошибке строго по формату из условия задачи.
func (e PositiveIntError) Error() string {
	//return "invalid number " + e.Input + ": "
	return fmt.Sprintf("invalid number %s: %s", e.Input, e.Problem)
}

func ParsePositiveInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		// TODO: Верните typed error PositiveIntError с заполненными полями (вместо "сырой" ошибки парсинга).
		return 0, PositiveIntError{Input: s, Problem: "not an int"}
	}

	if n <= 0 {
		// TODO: Верните typed error PositiveIntError с заполненными полями для случая неположительного числа.
		return 0, PositiveIntError{Input: s, Problem: "must be > 0"}
	}

	return n, nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan() // по условию читаем ровно одну строку
	s := sc.Text()

	n, err := ParsePositiveInt(s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(n)
}
