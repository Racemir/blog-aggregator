package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Racemir/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		fmt.Printf("usage: %v <name>\n", cmd.Name)
		return nil
	}

	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		if strings.Contains(err.Error(), "users_name_key") {
			return fmt.Errorf("user '%s' already exists", name)
		}
		return fmt.Errorf("couldn't create user: %w", err)
	}

	// Yeni kullanıcı oluşturulduysa config'e kaydet
	if err := s.cfg.SetUser(user.Name); err != nil {
		fmt.Printf("couldn't set current user: %v\n", err)
		return nil
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func handlerLogin(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %v", err)
	}

	if err := s.cfg.SetUser(name); err != nil {
		return fmt.Errorf("couldn't set current user: %v", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
