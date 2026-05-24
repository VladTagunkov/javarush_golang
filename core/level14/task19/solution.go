package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	// Читаем строку целиком до '\n' (если '\n' нет — возможен io.EOF, но строка может быть полезной).
	s, err := rd.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, "read error:", err)
		os.Exit(1)
	}

	found := false

	// Важно: range по строке идёт по рунам, а i — байтовый индекс.
	for i, r := range s {
		_ = i
		_ = r
		if r == '\t' || r == '\n' || r == '\r' {
			found = true
			fmt.Printf("pos=%d rune=%q code=%d\n", i, r, uint32(r))
		}
		// TODO: Найдите управляющие символы '\t', '\r', '\n' при обходе строки по рунам.
		// TODO: Для каждого найденного символа выведите отдельную строку в требуемом формате и отметьте found=true.
	}

	if !found {
		fmt.Println("clean")
	}
}
