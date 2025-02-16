package handlers

import (
	"context"
	"fmt"
	"github.com/grahamchill/blog_aggregator/internal"
	"github.com/grahamchill/blog_aggregator/internal/database"
)

func HandlerUnfollow(s *internal.State, cmd internal.Command, user database.User) error {
	// Ensure the correct number of arguments are provided
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: go run . unfollow <url>")
	}

	url := cmd.Args[0]

	// Delete the feed follow record
	err := s.Db.DeleteFeedFollowByUserAndURL(context.Background(), database.DeleteFeedFollowByUserAndURLParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return fmt.Errorf("failed to unfollow feed: %w", err)
	}

	// Print success message
	fmt.Printf("You have unfollowed the feed with URL: %s\n", url)
	return nil
}
