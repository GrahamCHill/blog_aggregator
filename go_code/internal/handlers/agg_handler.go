package handlers

import (
    "context"
    "fmt"
    "github.com/grahamchill/blog_aggregator/internal"
    "github.com/grahamchill/blog_aggregator/internal/RSS"
)

func HandlerAgg(s *internal.State, cmd internal.Command) error {
    // Ensure no additional arguments are provided
    if len(cmd.Args) > 0 {
        return fmt.Errorf("usage: go run . agg")
    }

    // Fetch the feed
    feed, err := RSS.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
    if err != nil {
        return fmt.Errorf("failed to fetch feed: %w", err)
    }

    // Print the entire struct
    fmt.Printf("%+v\n", feed)
    return nil
}
