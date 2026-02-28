package utils

import (
	"strconv"
)

func StringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}
