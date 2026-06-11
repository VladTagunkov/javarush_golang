package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TextFormatter interface {
	Format(s string) string
}

type PlainFormatter struct{}

func (PlainFormatter) Format(s string) string { return s }

type QuotedFormatter struct{}

func (QuotedFormatter) Format(s string) string {
	// TODO: Реализуйте форматирование строки в Go-стиле: результат должен быть в кавычках и с экранированием спецсимволов.
	return fmt.Sprintf("%q", s)
}

func Render(f TextFormatter, s string) string {
	// TODO: Реализуйте рендеринг строго через интерфейс: функция должна возвращать результат вызова f.Format(s).
	return f.Format(s)
}

func readLine(r *bufio.Reader) string {
	line, _ := r.ReadString('\n')
	return strings.TrimRight(line, "\r\n")
}

func main() {
	r := bufio.NewReader(os.Stdin)

	mode := readLine(r)
	text := readLine(r)

	var f TextFormatter
	if mode == "plain" {
		f = PlainFormatter{}
	} else if mode == "quoted" {
		f = QuotedFormatter{}
	} else {
		fmt.Println("unknown mode")
		return
	}

	fmt.Println(Render(f, text))
}
