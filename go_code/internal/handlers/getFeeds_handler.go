package handlers

import (
    "context"
    "fmt"
    "github.com/grahamchill/blog_aggregator/internal"
)

func HandlerFeeds(s *internal.State, cmd internal.Command) error {
    if len(cmd.Args) > 0 {
        return fmt.Errorf("usage: go run . feeds")
    }

    feeds, err := s.Db.GetFeeds(context.Background())
    if err != nil {
        return err
    }

    fmt.Printf("Feeds: \n")
    for _, feed := range feeds {
        fmt.Printf(" %+v\n", feed.Name)
        fmt.Printf(" %+v\n", feed.Url)
        fmt.Printf(" %+v\n\n", feed.UserName)
    }

    return nil
}
