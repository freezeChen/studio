/*
   @Time : 2019-05-08 17:13
   @Author : frozenchen
   @File : main
   @Software: studio
*/
package main

import (
	New "github.com/freezeChen/studio/new"
	"github.com/freezeChen/studio/template"
	"github.com/micro/cli"
	"github.com/micro/go-micro/cmd"
)

const version string = "v0.0.1"
const description = "a scaffolding for go-micro"

func main() {
	app := cmd.App()
	app.Commands = append(app.Commands, New.Command()...)

	app.Action = func(c *cli.Context) {
		cli.HelpPrinter(c.App.Writer, template.Help_template, c.App)
		cli.ShowAppHelp(c)
	}

	cmd.Init(
		cmd.Name("studio"),
		cmd.Description(description),
		cmd.Version(version))
}
