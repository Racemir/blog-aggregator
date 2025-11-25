package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset database: %v", err)
	}
	fmt.Printf("Database reset successfully\n")
	return nil
}
