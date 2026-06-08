package main

import "fmt"

type PartA struct {
	ID int
}

type PartB struct {
	ID int
}

// TODO: Объявите структуру Box так, чтобы она встраивала (embedding) и PartA, и PartB.
type Box struct {
	PartA
	PartB
}

func main() {
	var aID, bID int
	fmt.Scan(&aID, &bID)

	var box Box

	// TODO: Установите ID для обеих частей через явный путь: box.PartA.ID и box.PartB.ID.
	box.PartA.ID = aID
	box.PartB.ID = bID
	// TODO: Выведите ровно одну строку в формате A=<aID> B=<bID> (через явный доступ к полям).
	fmt.Printf("A=%d B=%d\n", box.PartA.ID, box.PartB.ID)

}
