package main

import "fmt"

func main() {
	nums := make([]int, 2, 5)

	// TODO: Заполните nums[0] и nums[1] значениями 7 и 8.
	nums[0] = 7
	nums[1] = 8

	fmt.Printf("len=%d cap=%d data=%v\n", len(nums), cap(nums), nums)

	i := 2

	// TODO: Реализуйте проверку доступа к nums[i] через if/else.
	// TODO: Условие допустимости должно быть строго i < len(nums) (не используйте cap).
	// TODO: Если индекс недопустим — выведите ровно "cannot access index 2", иначе — выведите nums[i].
	if i < len(nums) {
		fmt.Println(nums[i])
	} else {
		fmt.Println("cannot access index 2")
	}

}
