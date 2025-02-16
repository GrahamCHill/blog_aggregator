package RSS

import (
	"context"
	"fmt"
	"github.com/grahamchill/blog_aggregator/internal"
)

func ScrapeFeeds(s *internal.State) {
	ctx := context.Background()

	// Fetch the next feed to scrape
	feed, err := s.Db.GetNextFeedToFetch(ctx)
	if err != nil {
		fmt.Printf("Error fetching next feed: %v\n", err)
		return
	}

	// Mark the feed as fetched
	if err := s.Db.MarkFeedFetched(ctx, feed.ID); err != nil {
		fmt.Printf("Error marking feed as fetched: %v\n", err)
		return
	}

	// Fetch the RSS feed data
	rssFeed, err := FetchFeed(ctx, feed.Url)
	if err != nil {
		fmt.Printf("Error fetching feed %s: %v\n", feed.Url, err)
		return
	}

	// Print the titles of the feed's items
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("Feed: %s - Post: %s\n", feed.Url, item.Title)
	}
}
