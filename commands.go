package main

import (
	"errors"
)

func (c *commands) run(s *state, cmd command) error {
	function, ok := c.list[cmd.name]
	if !ok {
		return errors.New("command not found")
	}
	err := function(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.list[name] = f
}
