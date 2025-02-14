package main

import (
	"fmt"
	"github.com/grahamchill/blog_aggregator/internal"
	"github.com/grahamchill/blog_aggregator/internal/config"
	"os"
)

func main() {
	// Read the config file
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return
	}
	state := &internal.State{Cfg: cfg}

	cmds := &internal.Commands{
		Handlers: make(map[string]func(*internal.State, internal.Command) error),
	}

	cmds.Register("login", internal.HandlerLogin)

	if len(os.Args) < 2 {
		fmt.Printf("Error: Not enough arguments provided")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := internal.Command{
		Name: cmdName,
		Args: cmdArgs,
	}

	if err := cmds.Run(state, cmd); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
