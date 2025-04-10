package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	output := strings.Fields(lowerCase)

	return output

}
