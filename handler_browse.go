package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Racemir/blog-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	limit := 2
	if len(cmd.Args) == 1 {
		n, err := strconv.Atoi(cmd.Args[0])
		if err == nil {
			limit = n
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		Name:  user.Name,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts: %v", err)
	}

	for _, post := range posts {
		fmt.Printf("%s - %s (%s)\n", post.FeedsName, post.Title, post.Url)
	}

	return nil
}
