package main

import "fmt"

func main() {
	var packageCode int
	fmt.Scan(&packageCode)

	digitsSum := 0

	for packageCode > 0 {
		// TODO: Реализуйте вычисление суммы цифр числа.
		// TODO: На каждой итерации нужно получить последнюю цифру (packageCode % 10) и добавить её в digitsSum.
		// TODO: Затем нужно уменьшить packageCode (packageCode /= 10), чтобы перейти к следующей цифре.
		// TODO: Уберите break после того, как реализуете корректное обновление packageCode, иначе цикл выполнится только один раз.
		//break
		digitsSum += packageCode % 10
		packageCode /= 10
	}

	fmt.Print(digitsSum)
}
