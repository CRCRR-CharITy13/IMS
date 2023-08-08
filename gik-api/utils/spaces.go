package utils

import "strings"

func StandardizeSpaces(s string) (result string) {
	result = strings.Join(strings.Fields(s), " ")
	return
}
