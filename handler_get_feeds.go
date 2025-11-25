package main

import (
	"context"
	"fmt"

	"github.com/Racemir/blog-aggregator/internal/database"
)

func handlerGetFeeds(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %v", err)
	}
	for _, f := range feeds {
		fmt.Printf("* %s\n* %s\n", f.FeedName, f.UserName)
	}
	return nil
}
