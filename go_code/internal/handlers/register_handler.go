package handlers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/grahamchill/blog_aggregator/internal"
	"github.com/grahamchill/blog_aggregator/internal/config"
	"github.com/grahamchill/blog_aggregator/internal/database"
	"strings"
	"time"
)

func HandlerRegister(s *internal.State, cmd internal.Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: go run . register <name>")
	}
	name := cmd.Args[0]

	// Create a new user
	id := uuid.New()
	now := time.Now()
	user, err := s.Db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return fmt.Errorf("user with name '%s' already exists", name)
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Update the current user in the config
	s.Cfg.DbUser = name
	if err := config.Write(s.Cfg); err != nil {
		return fmt.Errorf("failed to update config: %w", err)
	}

	// Print success message
	fmt.Printf("User created: %+v\n", user)
	return nil
}
