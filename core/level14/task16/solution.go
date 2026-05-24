package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Заданный "сырой" вход: байты, разделитель внутри '|', пробелы и пустые элементы возможны.
	raw := []byte("  alpha| beta | |gamma|  ")

	// Чистим общий ввод и режем по байтовому разделителю.
	trimmed := bytes.TrimSpace(raw)
	parts := bytes.Split(trimmed, []byte("|"))

	// TODO: Пройдитесь по parts, у каждого элемента сделайте bytes.TrimSpace,
	// TODO: удалите элементы, которые стали пустыми, и соберите итоговый слайс clean.
	clean := make([][]byte, 0, len(parts))
	for _, part := range parts {
		clen_bytes := bytes.TrimSpace(part)
		if len(clen_bytes) != 0 {
			clean = append(clean, clen_bytes)
		}

	}

	// TODO: Соберите сводную байтовую строку line из clean через bytes.Join с разделителем []byte(",").
	var line []byte
	line = bytes.Join(clean, []byte(","))

	// Собираем многострочный отчёт через bytes.Buffer.
	var buf bytes.Buffer
	buf.WriteString("Items:\n")

	for _, item := range clean {
		buf.WriteString("- ")
		// TODO: Добавьте байты item и перевод строки '\n' через методы bytes.Buffer (не через конкатенацию строк).
		//_ = item
		buf.Write(item)
		buf.WriteString("\n")
	}

	buf.WriteString("Line: ")
	buf.Write(line)
	buf.WriteByte('\n')

	// Единичный финальный вывод.
	fmt.Print(buf.String())
}
