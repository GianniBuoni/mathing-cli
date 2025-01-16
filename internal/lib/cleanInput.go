package lib

import (
	"fmt"
	"strings"
)

func cleanInput(text string) ([]string, error) {
	if text == "" {
		return nil, fmt.Errorf("input string cannot be empty")
	}
	return strings.Fields(strings.ToLower(strings.TrimSpace(text))), nil
}
