package handlers

import (
	"context"
	"fmt"
	"github.com/grahamchill/blog_aggregator/internal"
	"github.com/grahamchill/blog_aggregator/internal/database"
)

func HandlerFollow(s *internal.State, cmd internal.Command) error {
	// Ensure the correct number of arguments are provided
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: go run . follow <url>")
	}

	url := cmd.Args[0]

	// Get the current user from the database
	user, err := s.Db.GetUser(context.Background(), s.Cfg.DbUser)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	// Look up the feed by URL
	feed, err := s.Db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to find feed: %w", err)
	}

	// Create the feed follow record
	feedFollow, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	// Print the result
	fmt.Printf("You are now following '%s' as %s.\n", feedFollow.FeedName, feedFollow.UserName)
	return nil
}
