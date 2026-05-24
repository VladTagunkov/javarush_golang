package main

import "fmt"

func main() {
	var codesCount int
	fmt.Scan(&codesCount)

	evenCodesCount := 0 // счётчик чётных кодов (zero value = 0)

	for i := 0; i < codesCount; i++ { // один цикл ровно на codesCount итераций
		var x int
		fmt.Scan(&x)

		// TODO: Определи, является ли код x чётным (через остаток от деления) и увеличь evenCodesCount на 1, если код чётный.
		if x%2 == 0 {
			evenCodesCount++
		}
	}

	fmt.Println(evenCodesCount)
}
