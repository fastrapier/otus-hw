package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(in string) (string, error) {
	var builder strings.Builder
	runes := []rune(in)

	// Пустая строка
	if len(runes) == 0 {
		return "", nil
	}

	for i := 0; i < len(runes); {
		var current rune

		// Если текущий символ – escape-символ
		if runes[i] == '\\' {
			// Если после обратного слэша нет символа, то это ошибка
			if i+1 >= len(runes) {
				return "", ErrInvalidString
			}

			next := runes[i+1]
			// Допускаются только цифры и обратный слэш
			if !unicode.IsDigit(next) && next != '\\' {
				return "", ErrInvalidString
			}

			current = next
			i += 2
		} else {
			// Если символ является цифрой и не экранирован, то это ошибка
			if unicode.IsDigit(runes[i]) {
				return "", ErrInvalidString
			}
			current = runes[i]
			i++
		}

		// Определение количества повторений
		rep := 1
		if i < len(runes) && unicode.IsDigit(runes[i]) {
			// Преобразуем цифру к числу
			num, err := strconv.Atoi(string(runes[i]))
			if err != nil {
				return "", ErrInvalidString
			}
			rep = num
			i++
		}

		// Добавляем символ rep раз, если rep не равно 0
		if rep > 0 {
			builder.WriteString(strings.Repeat(string(current), rep))
		}
	}
	return builder.String(), nil
}
