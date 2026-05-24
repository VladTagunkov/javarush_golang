package main

import "fmt"

// FilePerm — пользовательский тип для работы с битовой маской прав.
type FilePerm uint

const (
	// TODO: Объявите флаги прав через 1 << iota в строгом порядке: read, write, exec, hidden.
	PermRead FilePerm = 1 << iota
	PermWrite
	PermExec
	PermHidden
)

// hasPerm проверяет, установлен ли флаг p в маске all.
func hasPerm(all, p FilePerm) bool {
	// TODO: Реализуйте проверку установленного флага p в маске all через побитовые операции.
	// Важно: неизвестные биты в all не должны ломать результат.
	if all&p != 0 {
		return true
	}
	return false
}

func main() {
	var permCode uint
	fmt.Scan(&permCode)

	// Все проверки делаем через FilePerm.
	all := FilePerm(permCode)

	fmt.Printf("read=%t\n", hasPerm(all, PermRead))
	fmt.Printf("write=%t\n", hasPerm(all, PermWrite))
	fmt.Printf("exec=%t\n", hasPerm(all, PermExec))
	fmt.Printf("hidden=%t\n", hasPerm(all, PermHidden))
}
