package main

import (
	"github.com/mitchellh/cli"
)

type UserShow struct {
}

func UserShowCommand() (cli.Command, error) {
	return &UserShow{}, nil
}

func (c *UserShow) Help() string {
	return defaultHelpText
}

func (c *UserShow) Synopsis() string {
	return "Get details about an existing user"
}

func (c *UserShow) Run(args []string) int {
	return 0
}
