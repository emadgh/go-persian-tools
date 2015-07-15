package persiantools

import "strings"

var (
	englishNumbers []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	farsiNumbers   []string = []string{"۱", "۲", "۳", "۴", "۵", "۶", "۷", "۸", "۹", "۰"}
)

func NumbersToPersian(number string) string {
	if number == "" {
		return ""
	}
	for i, n := range englishNumbers {
		number = strings.Replace(number, n, farsiNumbers[i], -1)
	}
	return number
}

func NumbersToEnglish(number string) string {
	if number == "" {
		return ""
	}
	for i, n := range farsiNumbers {
		number = strings.Replace(number, n, englishNumbers[i], -1)
	}
	return number
}
