package internal

import "fmt"

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no arguments passed")
	}

	username := cmd.Args[0]
	err := s.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set user: %v", err)
	}

	fmt.Printf("User set to: %s\n", username)

	return nil
}
