package main

import "fmt"

func main() {
	var shortCount int
	var longCount int64

	fmt.Scan(&shortCount, &longCount)

	var sum int64
	// TODO: Посчитайте сумму в int64, явно приведя shortCount (int) к int64 перед сложением.
	sum = int64(shortCount) + longCount

	fmt.Println(sum)
	fmt.Printf("%T\n", sum)
}
