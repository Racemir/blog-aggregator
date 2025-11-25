package main

import (
	"context"
	"fmt"

	"github.com/Racemir/blog-aggregator/internal/database"
)

func handlerUsers(s *state, cmd command, user database.User) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get users: %v", err)
	}

	current := s.cfg.CurrentUserName
	for _, u := range users {
		if u.Name == current {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s\n", u.Name)
		}
	}
	return nil
}
