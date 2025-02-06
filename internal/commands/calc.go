package commands

import (
	"context"
	"fmt"
)

var calc CommandData = CommandData{
	Name:        "calc",
	Description: "Calculates totals per user in current receipt table.",
  Handler: HandleCalc,
}

func HandleCalc(s *State, cmd Command) error {
	ctx := context.Background()
	totals := map[int64]float64{}

	rowTotals, err := s.Store.GetRowTotal(ctx)
	if err != nil {
		return fmt.Errorf("issue getting row calculations %w", err)
	}
	for _, rowTotal := range rowTotals {
		price, ok := rowTotal.Total.(float64)
		if !ok {
			return fmt.Errorf("issue parsing row total, type mismatch: %v could not be cast as a float", rowTotal.Total)
		}
		if _, exists := totals[rowTotal.UserID]; !exists {
			totals[rowTotal.UserID] = price
			continue
		}
		totals[rowTotal.UserID] += price
	}
	for k, v := range totals {
		user, err := s.Store.GetUser(ctx, k)
		if err != nil {
			return fmt.Errorf("issue getting users %w", err)
		}
    fmt.Printf("%s: %05.2f\n", user.Name, v)
	}
	return nil
}
