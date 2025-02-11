package commands

import (
	"context"
	"fmt"
	"mathing/internal/store"
)

var calc CommandData = CommandData{
	Name:        "calc",
	Description: "Calculates totals per user in current receipt table.",
	Handler:     HandleCalc,
}

func HandleCalc(s *State, cmd Command) error {
	ctx := context.Background()
	totals := map[int64]float64{}

	rowTotals, err := s.Store.GetRowTotal(ctx)
	if err != nil {
		return fmt.Errorf("issue getting row calculations %w", err)
	}
	for _, rowTotal := range rowTotals {
		// parse if returned db value can be cast as float
		price, ok := rowTotal.Total.(float64)
		if !ok {
			return fmt.Errorf("issue parsing row total, type mismatch: %v could not be cast as a float", rowTotal.Total)
		}

		// parse payee string into slice of user ids
		users, err := store.PayeeIDToUserID(rowTotal.Payees)
		if err != nil {
			return err
		}

		for _, user := range users {
			if _, exists := totals[user]; !exists {
				totals[user] = price
				continue
			}
			totals[user] += price

		}

	}
	var grandTotal float64
	for k, v := range totals {
		user, err := s.Store.GetUser(ctx, k)
		if err != nil {
			return fmt.Errorf("issue getting users %w", err)
		}
		grandTotal += v
		fmt.Printf("%s: %05.2f\n", user.Name, v)
	}
	fmt.Printf("TOTAL: %05.2f\n", grandTotal)
	return nil
}
