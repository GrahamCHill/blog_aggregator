package internal

import "fmt"

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Handlers map[string]func(*State, Command) error
}

// Register - Registers a handler function to the Commands struct
func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Handlers[name] = f
}

// Run - If a command exists in the commands struct run it, else return an error
func (c *Commands) Run(s *State, cmd Command) error {
	handler, exists := c.Handlers[cmd.Name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}
