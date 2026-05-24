package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	tasks := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &tasks[i])
	}

	var k int
	fmt.Fscan(in, &k)

	m := 0
	m = min(k, n)
	// TODO: Вычислите фактический размер превью (m) как min(k, n).
	// TODO: Учтите, что k может быть больше n.

	preview := make([]string, 0)
	for i := 0; i < m; i++ {
		append_elem := strconv.Itoa(i+1) + ") " + tasks[i]
		preview = append(preview, append_elem)
	}
	// TODO: Соберите независимый слайс preview (не под-слайс tasks[:...]).
	// TODO: Нельзя изменять tasks.
	// TODO: Нельзя использовать copy(), slices.Clone и аналогичные клонирующие функции.
	// TODO: Используйте make + append и добавляйте строки в формате "1) <task>", "2) <task>" и т.д.

	//for i := 0; i < m; i++ {
	//	preview = append(preview, "")
	//}

	// Должно быть выведено ровно две строки: сначала preview, затем tasks.
	fmt.Println(preview)
	fmt.Println(tasks)
}
