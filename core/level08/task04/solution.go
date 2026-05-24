package main

import (
	"fmt"
)

func main() {
	var fileSize, blockSize int64
	fmt.Scan(&fileSize, &blockSize)

	// TODO: Проверьте входные данные и выведите INVALID, если fileSize < 0 или blockSize <= 0.
	// Сейчас проверка неполная — защищает только от деления на ноль.
	if (blockSize == 0) || (blockSize <= 0) || (fileSize < 0) {
		fmt.Println("INVALID")
		return
	}

	// TODO: Посчитайте количество блоков с округлением вверх, используя только целочисленную арифметику.
	// Текущая версия намеренно не учитывает "хвост" файла.
	blocks := (fileSize + blockSize - 1) / blockSize

	// Typed-константа (по заданию должна быть именно типизированной).
	const MaxBlocks uint8 = 255

	// TODO: Сравните blocks с MaxBlocks без ранней конвертации blocks к uint8.
	// Текущая версия намеренно неверная: при больших значениях произойдёт переполнение uint8.
	blocks8 := blocks

	if blocks8 > int64(MaxBlocks) {
		fmt.Println("TOO MANY")
		return
	} else {
		fmt.Println(uint8(blocks8))
	}

	// TODO: Делайте конвертацию к uint8 только в ветке, где blocks гарантированно не превышает MaxBlocks.

}
