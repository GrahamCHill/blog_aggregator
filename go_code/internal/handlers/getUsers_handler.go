package handlers

import (
    "context"
    "fmt"
    "github.com/grahamchill/blog_aggregator/internal"
)

func HandlerGetUsers(s *internal.State, cmd internal.Command) error {
    if len(cmd.Args) > 0 {
        return fmt.Errorf("usage: go run . users")
    }

    users, err := s.Db.GetUsers(context.Background())
    if err != nil {
        return err
    }

    for _, user := range users {
        if user == s.Cfg.DbUser {
            fmt.Printf("* %s (current)\n", user)
        } else {
            fmt.Printf("* %s\n", user)
        }
    }
    return nil
}
