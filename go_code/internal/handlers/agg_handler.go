package handlers

import (
    "fmt"
    "github.com/grahamchill/blog_aggregator/internal"
    "github.com/grahamchill/blog_aggregator/internal/RSS"
    "time"
)

func HandlerAgg(s *internal.State, cmd internal.Command) error {
    // Ensure one argument is provided
    if len(cmd.Args) != 1 {
        return fmt.Errorf("usage: go run . agg <time_between_reqs>")
    }

    // Parse the time duration
    timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
    if err != nil {
        return fmt.Errorf("invalid duration: %w", err)
    }

    fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

    // Create a time ticker
    ticker := time.NewTicker(timeBetweenRequests)
    defer ticker.Stop()

    // Run the scrapeFeeds function in a continuous loop
    for {
        RSS.ScrapeFeeds(s)
        <-ticker.C
    }
}
