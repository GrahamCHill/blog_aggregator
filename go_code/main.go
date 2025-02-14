package main

import (
	"database/sql"
	"fmt"
	"github.com/grahamchill/blog_aggregator/internal"
	"github.com/grahamchill/blog_aggregator/internal/config"
	"github.com/grahamchill/blog_aggregator/internal/database"
	"github.com/grahamchill/blog_aggregator/internal/handlers"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	// Read the config file
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return
	}

	// Open a database connection
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Printf("Error opening database connection: %v\n", err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// Initialize the database queries
	dbQueries := database.New(db)

	// Initialize the application state
	state := &internal.State{
		Db:  dbQueries,
		Cfg: cfg,
	}

	// Register command handlers
	cmds := &internal.Commands{
		Handlers: make(map[string]func(*internal.State, internal.Command) error),
	}
	cmds.Register("login", handlers.HandlerLogin)
	cmds.Register("register", handlers.HandlerRegister)

	// Check for command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Error: Not enough arguments provided")
		os.Exit(1)
	}

	// Parse the command and arguments
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := internal.Command{
		Name: cmdName,
		Args: cmdArgs,
	}

	// Execute the command
	if err := cmds.Run(state, cmd); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
