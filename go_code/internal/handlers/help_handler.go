package handlers

import (
	"fmt"
	"github.com/grahamchill/blog_aggregator/internal"
)

type HandlerFunc func(*internal.State, internal.Command) error

var handlers map[string]struct {
	Handler     HandlerFunc
	Description string
}

func init() {
	handlers = map[string]struct {
		Handler     HandlerFunc
		Description string
	}{
		"login": {
			Handler:     HandlerLogin,
			Description: "Log in as a user. Usage: go run . login <name>",
		},
		"register": {
			Handler:     HandlerRegister,
			Description: "Register a new user. Usage: go run . register <name>",
		},
		"addfeed": {
			Handler:     HandlerAddFeed,
			Description: "Add a new feed. Usage: go run . addfeed <name> <url>",
		},
		"reset": {
			Handler:     HandlerReset,
			Description: "Reset the database. Usage: go run . reset",
		},
		"users": {
			Handler:     HandlerGetUsers,
			Description: "List all users. Usage: go run . users",
		},
		"help": {
			Handler:     HandlerHelp,
			Description: "Display this help message. Usage: go run . help",
		},
		"agg": {
			Handler:     HandlerAgg,
			Description: "Display aggregated data of xml feed. Usage: go run . agg",
		},
	}
}

func HandlerHelp(s *internal.State, cmd internal.Command) error {
	fmt.Println("Available commands:")
	for name, h := range handlers {
		fmt.Printf("  %s: %s\n", name, h.Description)
	}
	return nil
}
