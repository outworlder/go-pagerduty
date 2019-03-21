package main

import (
	"github.com/mitchellh/cli"
)

type SchedulePreview struct {
}

func SchedulePreviewCommand() (cli.Command, error) {
	return &SchedulePreview{}, nil
}

func (c *SchedulePreview) Help() string {
	return defaultHelpText
}

func (c *SchedulePreview) Synopsis() string {
	return "Preview what an on-call schedule would look like without saving it"
}

func (c *SchedulePreview) Run(args []string) int {
	return 0
}
