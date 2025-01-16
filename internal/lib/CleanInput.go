package lib

import (
	"fmt"
	"strings"
)

func CleanInput(text string) ([]string, error) {
	if text == "" {
		return nil, fmt.Errorf("input string cannot be empty")
	}
	return strings.Fields(strings.ToLower(strings.TrimSpace(text))), nil
}
