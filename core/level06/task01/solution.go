package main

import "fmt"

func main() {
	var budget int
	var purchase int
	var cashback int

	fmt.Scan(&budget, &purchase, &cashback)

	rest := budget
	rest -= purchase
	rest += cashback
	// TODO: Учтите покупку: обновите rest через "=", не используйте ":=" повторно для rest
	// TODO: Учтите кешбэк: снова обновите rest через "=", не используйте ":=" повторно для rest

	fmt.Println(rest)
}
