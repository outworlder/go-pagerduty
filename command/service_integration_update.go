package main

import (
	"github.com/mitchellh/cli"
)

type ServiceIntegrationUpdate struct {
}

func ServiceIntegrationUpdateCommand() (cli.Command, error) {
	return &ServiceIntegrationUpdate{}, nil
}

func (c *ServiceIntegrationUpdate) Help() string {
	return defaultHelpText
}

func (c *ServiceIntegrationUpdate) Synopsis() string {
	return "Update an integration belonging to a service"
}

func (c *ServiceIntegrationUpdate) Run(args []string) int {
	return 0
}
