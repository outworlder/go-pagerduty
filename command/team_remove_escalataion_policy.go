package main

import (
	"github.com/mitchellh/cli"
)

type TeamRemoveEscalationPolicy struct {
}

func TeamRemoveEscalationPolicyCommand() (cli.Command, error) {
	return &TeamRemoveEscalationPolicy{}, nil
}

func (c *TeamRemoveEscalationPolicy) Help() string {
	return defaultHelpText
}

func (c *TeamRemoveEscalationPolicy) Synopsis() string {
	return "Remove an escalation policy from a team"
}

func (c *TeamRemoveEscalationPolicy) Run(args []string) int {
	return 0
}
