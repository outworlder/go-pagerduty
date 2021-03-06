package main

import (
	"github.com/mitchellh/cli"
)

type MaintenanceWindowDelete struct {
}

func MaintenanceWindowDeleteCommand() (cli.Command, error) {
	return &MaintenanceWindowDelete{}, nil
}

func (c *MaintenanceWindowDelete) Help() string {
	return defaultHelpText
}

func (c *MaintenanceWindowDelete) Synopsis() string {
	return "Delete an existing maintenance window if it's in the future, or end it if it's currently on-going"
}

func (c *MaintenanceWindowDelete) Run(args []string) int {
	return 0
}
