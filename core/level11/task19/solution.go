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

	// Сценарий без предвыделения: var a []int
	var a []int
	prevCapA := cap(a)
	noPrealloc := 0

	for i := 0; i < n; i++ {
		a = append(a, i)

		// TODO: Реализуйте подсчёт количества изменений ёмкости cap(a) относительно prevCapA.
		// TODO: Счётчик noPrealloc должен увеличиваться только когда cap(a) изменился.
		// TODO: Не забудьте обновлять prevCapA после фиксации изменения.
		if prevCapA != cap(a) {
			noPrealloc++
			prevCapA = cap(a)
		}
	}

	// Сценарий с предвыделением: b := make([]int, 0, n)
	b := make([]int, 0, n)
	prevCapB := cap(b)
	prealloc := 0

	for i := 0; i < n; i++ {
		b = append(b, i)

		// TODO: Реализуйте подсчёт количества изменений ёмкости cap(b) относительно prevCapB.
		// TODO: Счётчик prealloc должен увеличиваться только когда cap(b) изменился.
		// TODO: Не забудьте обновлять prevCapB после фиксации изменения.
		if prevCapB != cap(b) {
			prealloc++
			prevCapB = cap(b)
		}
	}

	fmt.Fprintf(out, "no_prealloc=%d prealloc=%d", noPrealloc, prealloc)
}
