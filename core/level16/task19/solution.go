package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseQuantity(s string) (int, error) {
	// TODO: Распарсить строку в целое число через strconv.Atoi; при ошибке вернуть (0, error) с корректным текстом ошибки
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("parse error")
	}
	// TODO: Если парсинг успешен — проверить диапазон 1..1000 включительно; при нарушении вернуть (0, error) с корректным текстом ошибки
	if val > 0 && val <= 1000 {
		return val, nil
	}

	return 0, errors.New("value out of range")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := strings.TrimSpace(scanner.Text())

	quantity, err := ParseQuantity(s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("OK %d\n", quantity)
}
