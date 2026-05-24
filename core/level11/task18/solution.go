package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	// TODO: Создайте слайс values строго через make([]int, 0, n), чтобы len был 0, а cap был n.
	values := make([]int, 0, n)

	sum := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)

		// TODO: Добавьте x в values только через append (без индексной записи) и посчитайте сумму.
		values = append(values, x)
		sum += x
	}

	// TODO: Выведите в первой строке все числа из values в исходном порядке через пробел.
	var res_str string
	for i, v := range values {
		if i != 0 {
			res_str += " " + strconv.Itoa(v)
		} else {
			res_str += strconv.Itoa(v)
		}
	}
	fmt.Fprintln(out, res_str)
	//fmt.Print(res_str)
	fmt.Fprintf(out, "len=%d cap=%d\n", len(values), cap(values))
	fmt.Fprintln(out, sum)
}
