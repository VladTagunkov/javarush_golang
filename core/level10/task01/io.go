package main

import "fmt"

// readTwoInts отвечает только за ввод: читает ровно два int и возвращает их.
func readTwoInts() (a int, b int) {
	// TODO: Прочитайте из stdin ровно два целых числа (int) и верните их.
	_, _ = fmt.Scan(&a, &b)
	return a, b
}
