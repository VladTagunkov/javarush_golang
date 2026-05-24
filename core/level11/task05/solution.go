package main

import "fmt"

func main() {
	// TODO: Создайте слайс nums через литерал с нужными значениями из условия.
	nums := []int{10, 20, 30}

	// TODO: Выведите len и cap в точном формате, который требуется в условии (используйте fmt.Printf).
	fmt.Printf("len=%d cap=%d\n", len(nums), cap(nums))

	// Вторая строка печатается обычным способом.
	fmt.Println(nums)
}
