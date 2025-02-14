package handlers

import (
	"context"
	"fmt"
	"github.com/grahamchill/blog_aggregator/internal"
	"github.com/grahamchill/blog_aggregator/internal/config"
	"strings"
)

func HandlerLogin(s *internal.State, cmd internal.Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: go run . login <name>")
	}
	name := cmd.Args[0]

	// Check if the user exists
	user, err := s.Db.GetUser(context.Background(), name)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return fmt.Errorf("user with name '%s' does not exist", name)
		}
		return fmt.Errorf("failed to get user: %w", err)
	}

	// Update the current user in the config
	s.Cfg.DbUser = user.Name
	if err := config.Write(s.Cfg); err != nil {
		return fmt.Errorf("failed to update config: %w", err)
	}

	// Print success message
	fmt.Printf("Logged in as: %s\n", user.Name)
	return nil
}
