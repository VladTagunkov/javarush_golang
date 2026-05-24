package main

import (
	"fmt"
	"sort"
)

func main() {
	freq := make(map[string]int)

	// Читаем слова до первой ошибки ввода (в т.ч. EOF).
	for {
		var w string
		if _, err := fmt.Scan(&w); err != nil {
			break
		}
		freq[w]++
	}

	// Собираем ключи map в отдельный слайс, чтобы не зависеть от порядка map.
	words := make([]string, 0, len(freq))
	for w := range freq {
		words = append(words, w)
	}

	// Временная (неполная) сортировка: только по алфавиту.
	// TODO: Реализуйте сортировку words по убыванию частоты, а при равенстве частоты — по возрастанию слова.
	sort.Slice(words, func(i, j int) bool {
		if freq[words[i]] != freq[words[j]] {
			return freq[words[i]] > freq[words[j]]
		}
		return words[i] < words[j]
	})

	// Печать строго по слайсу (не через range по map).
	for _, w := range words {
		fmt.Printf("%s %d\n", w, freq[w])
	}
}
