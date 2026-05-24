package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var _ = errors.Join

func ParseInts(values []string) ([]int, error) {
	// TODO: Реализуйте преобразование КАЖДОЙ строки values[i] в int строго через strconv.Atoi.
	// TODO: При ошибке парсинга не прерывайте цикл: добавляйте ошибку с контекстом индекса и исходной строки через fmt.Errorf("index %d (%q): %w", i, values[i], err).
	// TODO: Верните слайс только успешно распарсенных чисел (в исходном порядке) и одну ошибку, объединяющую все ошибки через errors.Join.

	ints := make([]int, 0, len(values))
	var errs_list []error

	for i, s := range values {
		n, err := strconv.Atoi(s)
		if err != nil {
			// Ошибки пока игнорируются — студент должен реализовать агрегацию ошибок.
			err_val := fmt.Errorf("index %d (%q): %w", i, values[i], err)
			errs_list = append(errs_list, err_val)
			continue
		}
		ints = append(ints, n)
	}

	return ints, errors.Join(errs_list...)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var m int
	fmt.Fscan(in, &m)

	values := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &values[i])
	}

	ints, err := ParseInts(values)

	sum := 0
	for _, n := range ints {
		sum += n
	}

	if err == nil {
		fmt.Println("OK")
		fmt.Printf("sum=%d\n", sum)
		return
	}

	fmt.Println(err)
	fmt.Printf("sum=%d\n", sum)
}
