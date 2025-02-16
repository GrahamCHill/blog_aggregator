package handlers

import (
	"context"
	"fmt"
	"github.com/grahamchill/blog_aggregator/internal"
)

func HandlerFollowing(s *internal.State, cmd internal.Command) error {
	// Ensure no additional arguments are provided
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: go run . following")
	}

	// Get the current user from the database
	user, err := s.Db.GetUser(context.Background(), s.Cfg.DbUser)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	// Fetch all feed follows for the user
	feedFollows, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to fetch feed follows: %w", err)
	}

	// Print the feeds the user is following
	fmt.Println("You are following these feeds:")
	for _, follow := range feedFollows {
		fmt.Printf("- %s\n", follow.FeedName)
	}

	return nil
}
