package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	const slash rune = '\\'
	var result strings.Builder
	strRune := []rune(str)

	fmt.Printf("%T %+v\n", strRune, strRune)
	if str == "" {
		return "", nil
	}
	// если первая руна число
	if unicode.IsDigit(strRune[0]) {
		return "", ErrInvalidString
	}

	for i, v := range strRune {
		var nextRune rune

		if len(strRune) > i+1 {
			nextRune = strRune[i+1]
		}

		// если текущая руна '\' и следующая не '\' пропускаем
		if v == slash && nextRune != slash {
			continue
		}

		if unicode.IsDigit(v) {
			prevRune := strRune[i-1]
			// если следующая руна тоже число выводим ошибку
			if unicode.IsDigit(nextRune) && prevRune != slash {
				return "", ErrInvalidString
			}

			// если предыдущая руна '\' тогда то просто записываем число в результат
			if prevRune == slash && strRune[i-2] != slash {
				result.WriteString(string(v))
				continue
			}

			repeatSymbol(&result, v, prevRune)
		} else {
			result.WriteString(string(v))
		}
	}

	return result.String(), nil
}

// Повторение руны и запись в результат указанное количество раз -1 (вы читаем ранее добавленную к строке руну).
func repeatSymbol(result *strings.Builder, quantity rune, prevRune rune) {
	count, _ := strconv.Atoi(string(quantity))
	// если ноль то удаляем предыдущий символ
	if count == 0 {
		prevResult := []rune(result.String())
		result.Reset()
		result.WriteString(string(prevResult[:len(prevResult)-1]))
	} else {
		result.WriteString(strings.Repeat(string(prevRune), count-1))
	}
}
