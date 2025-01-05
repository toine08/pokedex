package utils

import (
	"strings"
)

func CleanInput(text string) []string {
	var sliceText []string
	textStrings := strings.Fields(strings.ToLower(text))
	sliceText = append(sliceText, textStrings...)
	return sliceText
}
