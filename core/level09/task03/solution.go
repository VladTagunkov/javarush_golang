package main

import "fmt"

func max2(x, y int) int {
	// TODO: Реализуйте функцию, которая возвращает большее из двух целых чисел (при равенстве верните это же значение).
	if x > y {
		return x
	} else if y > x {
		return y
	} else if x == y {
		return x
	} else {
		return y
	}

}

func max3(a, b, c int) int {
	// TODO: Реализуйте максимум из трёх чисел через вызовы max2 (без одного большого условия, сравнивающего сразу три значения).
	// Временная логика: сравниваем только первые два значения, третье игнорируется.
	if max2(a, b) > c {
		return max2(a, b)
	} else if max2(b, c) > a {
		return max2(b, c)
	} else if max2(a, c) > b {
		return max2(a, c)
	} else {
		return max2(a, b)
	}
}

func main() {
	var firstScore, secondScore, thirdScore int
	fmt.Scan(&firstScore, &secondScore, &thirdScore)

	fmt.Print(max3(firstScore, secondScore, thirdScore))
}
