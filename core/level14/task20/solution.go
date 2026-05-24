package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	// По условию читаем ровно две строки именно через ReadString('\n').
	expectedRaw, _ := r.ReadString('\n')
	actualRaw, _ := r.ReadString('\n')

	// Norm строго через TrimSpace от raw-строк.
	expectedNorm := strings.TrimSpace(expectedRaw)
	actualNorm := strings.TrimSpace(actualRaw)

	if expectedNorm == actualNorm {
		fmt.Println("match")
		return
	}

	fmt.Printf("expectedRaw=<%q>\n", expectedRaw)
	fmt.Printf("actualRaw=<%q>\n", actualRaw)
	fmt.Printf("expectedNorm=<%q>\n", expectedNorm)
	fmt.Printf("actualNorm=<%q>\n", actualNorm)

	idx, expB, actB := firstDiff(expectedNorm, actualNorm)
	fmt.Printf("firstDiff=%d %d %d\n", idx, expB, actB)
}

func firstDiff(expectedNorm, actualNorm string) (idx int, expByte int, actByte int) {
	// TODO: Реализуйте поиск первого отличия на уровне байтов в нормализованных строках.
	// TODO: Нужно сравнивать именно []byte(expectedNorm) и []byte(actualNorm).
	// TODO: Верните индекс отличия, байт expected (или -1 если expected закончилась), байт actual (или -1 если actual закончилась).
	// TODO: Учтите случай, когда одна строка является префиксом другой (отличие только в длине).
	expn := []byte(expectedNorm)
	actn := []byte(actualNorm)

	//var idx int

	if len(expn) == len(actn) {
		for i, b := range expn {
			if i <= len(actn) || i <= len(expn) {
				if b != actn[i] {
					idx = i
					expByte = int(b)
					actByte = int(actn[i])
					break
				}
			} else {
				idx = -1
				expByte = -1
				actByte = -1
			}
		}
	} else if len(expn) > len(actn) {
		for i, b := range expn {
			if i <= len(actn) || i <= len(expn) {
				if b != actn[i] {
					idx = i
					expByte = int(b)
					actByte = int(actn[i])
					break
				}
			} else {
				idx = -1
				expByte = -1
				actByte = -1
			}
		}
	}
	return idx, expByte, actByte
}
