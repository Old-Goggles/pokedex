package main

import (
	"strings"
)

func cleanInput(text string) []string {
	trim := strings.TrimSpace(text)
	lower := strings.ToLower(trim)
	field := strings.Fields(lower)

	return field
}
