package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	words := strings.Split(text, " ")
	res := []string{}

	for _, w := range words {
		if w == "" {
			continue
		}

		res = append(res, w)

	}

	return res

}
