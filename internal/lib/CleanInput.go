package lib

import (
	"strings"
)

func CleanInput(text string) (string, error) {
	return strings.ToLower(strings.TrimSpace(text)), nil
}
