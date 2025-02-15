package handlers

import (
    "context"
    "fmt"
    "github.com/google/uuid"
    "github.com/grahamchill/blog_aggregator/internal"
    "github.com/grahamchill/blog_aggregator/internal/database"
)

func HandlerAddFeed(s *internal.State, cmd internal.Command, user database.User) error {
    // Ensure the correct number of arguments are provided
    if len(cmd.Args) != 2 {
        return fmt.Errorf("usage: go run . addfeed <name> <url>")
    }

    name := cmd.Args[0]
    url := cmd.Args[1]

    feedID := uuid.New()

    // Create the feed
    feed, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
        ID:     feedID, // Pass the generated UUID
        Name:   name,
        Url:    url,
        UserID: user.ID,
    })
    if err != nil {
        return fmt.Errorf("failed to create feed: %w", err)
    }

    // Automatically create a feed follow record
    _, err = s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
        UserID: user.ID,
        FeedID: feed.ID,
    })
    if err != nil {
        return fmt.Errorf("failed to create feed follow: %w", err)
    }

    // Print the new feed record
    fmt.Printf("Feed created successfully:\n")
    fmt.Printf("ID: %s\n", feed.ID)
    fmt.Printf("Name: %s\n", feed.Name)
    fmt.Printf("URL: %s\n", feed.Url)
    fmt.Printf("User ID: %s\n", feed.UserID)
    fmt.Printf("Created At: %s\n", feed.CreatedAt)
    fmt.Printf("Updated At: %s\n", feed.UpdatedAt)

    return nil
}
