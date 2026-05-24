package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	s := ""
	var err error
	s, err = r.ReadString('\n')

	// TODO: Считайте ровно одну строку из stdin через bufio.Reader так, чтобы строка включала завершающий '\n' (если он есть во входе).
	// TODO: Обработайте ошибку чтения без паники: при ошибке выведите понятное сообщение и завершите программу.

	if err != nil {
		fmt.Println("We have error:", err)
		return
	}

	fmt.Printf("raw=%q\n", s)
	fmt.Printf("len=%d\n", len(s))
	fmt.Printf("bytes=%v\n", []byte(s))
}
