package handlers

import (
	"context"
	"fmt"
	"github.com/grahamchill/blog_aggregator/internal"
)

func HandlerReset(s *internal.State, cmd internal.Command) error {
	// Ensure no additional arguments are provided
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: go run . reset")
	}

	// Call the DeleteAllUsers query
	err := s.Db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to delete all users: %w", err)
	}

	// Print success message
	fmt.Println("All users deleted successfully")
	return nil
}
