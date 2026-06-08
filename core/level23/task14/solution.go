package main

import "fmt"

type Meta struct {
	Version int
}

type Document struct {
	Meta
	// TODO: Добавьте собственное поле Version типа int, чтобы оно перекрывало Meta.Version при обращении d.Version.
	Version int
}

func main() {
	var docVersion, metaVersion int
	fmt.Scan(&docVersion, &metaVersion)

	d := Document{}

	// TODO: Инициализируйте Document так, чтобы d.Version было равно docVersion, а d.Meta.Version было равно metaVersion.
	d.Meta.Version = metaVersion
	d.Version = docVersion

	// TODO: Выведите строку в нужном формате и посчитайте diff как разницу версии документа и версии в мета-информации.
	// TODO: Для версии внутри мета-информации используйте явный доступ d.Meta.Version.
	fmt.Printf("doc=%d meta=%d diff=%d", d.Version, d.Meta.Version, d.Version-d.Meta.Version)
}
