/*
   @Time : 2019-05-08 17:13
   @Author : frozenchen
   @File : main
   @Software: studio
*/
package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro/cmd"
	New "studio/new"
)

const version string = "v0.0.1"
const description = "a scaffolding for go-micro"

func main() {

	app := cmd.App()
	app.Flags = append(app.Flags,
		cli.StringFlag{
			Name:  "test",
			Usage: "fdsfdsf",
		})

	//cli.hel

	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	app.Commands = append(app.Commands, New.Command()...)

	cmd.Init(
		cmd.Name("studio"),
		cmd.Description(description),
		cmd.Version(version))

}

func setup() {

}
