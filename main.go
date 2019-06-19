package main

import (
	"github.com/freezeChen/studio/new"
	"github.com/urfave/cli"
	"os"
)

const (
	Version = "v0.0.7"
)

func main() {
	app := cli.NewApp()
	app.Name = "studio"
	app.Usage = "studio工具集"
	app.Version = Version

	command := new.Command()
	command = append(command,
		cli.Command{
			Name:            "tool",
			Aliases:         []string{"t"},
			Usage:           "studio tool",
			Action:          toolAction,
			SkipFlagParsing: true,
		},
	)
	app.Commands = command

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
