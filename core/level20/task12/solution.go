package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var name, ageStr, email string
	fmt.Fscan(in, &name, &ageStr, &email)

	if err := ValidateProfile(name, ageStr, email); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print("ok")
}

func ValidateProfile(name, ageStr, email string) error {
	var errs []error

	// TODO: Реализуйте валидацию так, чтобы проверки не останавливались на первой ошибке,
	// TODO: а добавляли все найденные проблемы в errs в стабильном порядке: name -> age -> email.

	if strings.TrimSpace(name) == "" {
		errs = append(errs, errors.New("name is empty"))
	}
	if strings.TrimSpace(ageStr) == "" {
		errs = append(errs, errors.New("age is empty"))
	}
	if strings.TrimSpace(email) == "" {
		errs = append(errs, errors.New("email is empty"))
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		errs = append(errs, errors.New("age is not a number"))
	} else {
		//_ = age
		// TODO: Добавьте проверку диапазона возраста [0..150] и отдельную ошибку для выхода за диапазон.
		if age < 0 || age > 150 {
			errs = append(errs, errors.New("age is out of range"))
		}
	}

	if !strings.Contains(email, "@") {
		errs = append(errs, errors.New("email must contain @"))
	} else {
		// TODO: Добавьте проверку, что после @ есть хотя бы одна точка '.' (это отдельная ошибка).
		if !strings.Contains(email, "@.") {
			errs = append(errs, errors.New("no @ symbol"))
		}
	}

	if len(errs) == 0 {
		return nil
	}

	// TODO: Верните одну агрегированную ошибку через errors.Join(errs...), а не только первую ошибку.
	return errors.Join(errs...)
}
