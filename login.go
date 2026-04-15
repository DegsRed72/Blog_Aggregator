package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("login expects only one argument")
	}
	s.config.SetUser(cmd.args[0])
	fmt.Println("Username has been set")
	return nil
}
