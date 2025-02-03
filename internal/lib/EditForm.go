package lib

import (
	"fmt"
	"mathing/internal/store"
)

func WithStartValue(s store.Item) func(*ItemFormData) {
	return func(i *ItemFormData) {
		i.Item = s.Item
		i.Price = fmt.Sprintf("%05.2f", s.Price)
	}
}
