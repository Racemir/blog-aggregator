package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Racemir/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't current user: %v", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed %v", err)
	}

	if _, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}); err != nil {
		return fmt.Errorf("couldn't create fallow for new feed: %v", err)
	}

	fmt.Println("Feed added successfully:")
	fmt.Printf(" * ID: %v\n", feed.ID)
	fmt.Printf(" * Name: %v\n", feed.Name)
	fmt.Printf(" * URL: %v\n", feed.Url)
	fmt.Printf(" * UserID: %v\n", feed.UserID)

	return nil
}
