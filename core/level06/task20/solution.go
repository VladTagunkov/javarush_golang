package main

import "fmt"

var appName string = "MiniApp"
var debug bool // по умолчанию false

func main() {
	var flag int
	fmt.Scan(&flag)

	var mode string

	// TODO: Установите глобальную переменную debug и строку mode в зависимости от значения flag (0 или 1).
	// TODO: Важно: не создавайте новую локальную переменную debug через ":=" — нужно обновлять глобальную через "=".
	if flag == 1 {
		debug = true
		mode = "debug"
	} else if flag == 0 {
		debug = false
		mode = "release"
	}

	fmt.Printf("%s mode=%s\n", appName, mode)
}
