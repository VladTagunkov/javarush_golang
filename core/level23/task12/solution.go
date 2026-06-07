package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Saver — базовая операция записи.
type Saver interface {
	Save(key, value string) error
}

// Loader — базовая операция чтения.
type Loader interface {
	Load(key string) (string, bool)
}

// Deleter — базовая операция удаления.
type Deleter interface {
	Delete(key string) error
}

// KeyValueStore собирает базовые контракты через embedding.
type KeyValueStore interface {
	Saver
	Loader
	Deleter
}

// Dumper — административная возможность: стабильный дамп.
type Dumper interface {
	Dump() []string
}

// AdminStore — расширенный контракт: базовые операции + дамп.
type AdminStore interface {
	KeyValueStore
	Dumper
}

// MapStore — простая реализация на базе map[string]string.
type MapStore struct {
	data map[string]string
}

func (s *MapStore) Save(key, value string) error {
	// map может быть nil у zero-value — создаём при первом сохранении
	if s.data == nil {
		s.data = make(map[string]string)
	}
	s.data[key] = value
	return nil
}

func (s *MapStore) Load(key string) (string, bool) {
	v, ok := s.data[key]
	return v, ok
}

func (s *MapStore) Delete(key string) error {
	// TODO: Реализуйте удаление ключа из хранилища. Удаление отсутствующего ключа не должно приводить к ошибке.
	for k := range s.data {
		if k == key {
			delete(s.data, k)
			return nil
		}
	}
	return nil
}

func (s *MapStore) Dump() []string {
	// TODO: Реализуйте дамп хранилища: верните []string в формате key=value, отсортированный по key по возрастанию.
	var keys []string
	for k := range s.data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var result []string
	for _, k := range keys {
		result = append(result, k+"="+s.data[k])
	}
	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	var store MapStore
	for i := 0; i < n; i++ {
		var key, value string
		fmt.Fscan(in, &key, &value)
		store.Save(key, value)
	}

	var delKey string
	fmt.Fscan(in, &delKey)
	store.Delete(delKey)

	for _, line := range store.Dump() {
		fmt.Fprintln(out, line)
	}
}
