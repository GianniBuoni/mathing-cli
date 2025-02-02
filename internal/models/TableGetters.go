package models

import (
	"math"
)

func (t *TableData) CurrentPage() int64 {
	return (t.pageOffset / 20) + 1
}

func (t *TableData) PageCount() int64 {
	return int64(math.Ceil(float64(t.itemCount) / 20))
}
