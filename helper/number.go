package helper

import (
	"strconv"
)

func NumberOfDigits(number int) int {
	if number < 10 {
		return 1
	}

	return 1 + NumberOfDigits(number/10)
}

func ToString(number int, digits ...int) string {
	str := strconv.Itoa(number)
	if len(digits) == 0 || digits[0] <= 0 {
		return str
	}

	digit := digits[0]
	for i := len(str); i < digit; i++ {
		str = "0" + str
	}

	return str
}
