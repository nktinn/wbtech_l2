package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// UnpackString распаковывает строку
func UnpackString(s string) (string, error) {
	var result bytes.Buffer
	for i, r := range s {
		if unicode.IsDigit(r) {
			if i == 0 || unicode.IsDigit(rune(s[i-1])) {
				return "", errors.New("invalid string")
			}
			count, _ := strconv.Atoi(string(r))

			for j := 0; j < count-1; j++ {
				result.WriteRune(rune(s[i-1]))
			}
		} else {
			result.WriteRune(r)
		}
	}
	return result.String(), nil
}

func main() {
	str, err := UnpackString("a4bc2d5e")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}
