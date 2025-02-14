package main

import (
	"fmt"
	"github.com/grahamchill/blog_aggregator/internal/config"
)

func main() {
	// Read the config file
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return
	}

	// Set the current user to your name (e.g., "lane")
	err = cfg.SetUser("Graham")
	if err != nil {
		fmt.Printf("Error setting user: %v\n", err)
		return
	}

	// Read the config file again
	updatedCfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading updated config: %v\n", err)
		return
	}

	// Print the contents of the config struct
	fmt.Printf("Updated Config: %+v\n", updatedCfg)
}
