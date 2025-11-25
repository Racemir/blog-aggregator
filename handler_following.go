package main

import (
	"context"
	"fmt"

	"github.com/Racemir/blog-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %v", err)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %v", err)
	}

	for _, f := range follows {
		fmt.Println(f.FeedName)
	}
	return nil
}
