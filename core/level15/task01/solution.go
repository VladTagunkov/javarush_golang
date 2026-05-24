// package main
//
// import "fmt"
//
//	func main() {
//		// Мини-словарь "английское слово -> перевод" (обязательно литералом, без make).
//		dictionary := map[string]string{
//			"hello":  "привет",  // TODO: Заполните перевод для слова "hello".
//			"bye":    "пока",    // TODO: Заполните перевод для слова "bye".
//			"thanks": "спасибо", // TODO: Заполните перевод для слова "thanks".
//		}
//
//		var inputWord string
//		// TODO: Считайте одно английское слово из stdin в переменную inputWord с помощью fmt.Scan или fmt.Fscan.
//		fmt.Scan(&inputWord)
//		// По условию слово всегда есть в словаре — ok-проверка запрещена/не нужна.
//		fmt.Println(dictionary[inputWord])
//		fmt.Println(len(dictionary))
//	}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	sums := make(map[string]int)

	for i := 0; i < n; i++ {
		var category string
		var amount int
		fmt.Fscan(in, &category, &amount)

		// TODO: Реализуйте корректное накопление суммы по категориям (повторы должны увеличивать сумму, а не затирать предыдущее значение).

		sums[category] += amount
	}

	total := 0
	for _, v := range sums {
		total += v
	}

	fmt.Fprintln(out, total)

}
