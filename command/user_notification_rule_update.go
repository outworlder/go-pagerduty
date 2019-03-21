package main

import (
	"github.com/mitchellh/cli"
)

type UserNotificationRuleUpdate struct {
}

func UserNotificationRuleUpdateCommand() (cli.Command, error) {
	return &UserNotificationRuleUpdate{}, nil
}

func (c *UserNotificationRuleUpdate) Help() string {
	return defaultHelpText
}

func (c *UserNotificationRuleUpdate) Synopsis() string {
	return "Update a user's notification rule"
}

func (c *UserNotificationRuleUpdate) Run(args []string) int {
	return 0
}
