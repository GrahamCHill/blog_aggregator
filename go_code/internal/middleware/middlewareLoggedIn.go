package middleware

import (
    "context"
    "fmt"
    "github.com/grahamchill/blog_aggregator/internal"
    "github.com/grahamchill/blog_aggregator/internal/database"
)

// middlewareLoggedIn wraps a handler that requires a logged-in user.
func MiddlewareLoggedIn(handler func(s *internal.State, cmd internal.Command, user database.User) error) func(*internal.State, internal.Command) error {
    return func(s *internal.State, cmd internal.Command) error {
        // Get the current user from the database
        user, err := s.Db.GetUser(context.Background(), s.Cfg.DbUser)
        if err != nil {
            return fmt.Errorf("user not logged in: %w", err)
        }

        // Call the wrapped handler with the user
        return handler(s, cmd, user)
    }
}
