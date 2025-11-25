package main

import (
	"fmt"
	"time"

	"github.com/Racemir/blog-aggregator/internal/database"
)

func handlerAgg(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs", cmd.Name)
	}
	timeBetweenReauests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %v", err)
	}
	fmt.Printf("Collecting feeds every %v\n", timeBetweenReauests)

	for {
		err := scrapeFeeds(s)
		if err != nil {
			fmt.Println("Error scraping feeds:", err)
		}
		time.Sleep(timeBetweenReauests)
	}
}
