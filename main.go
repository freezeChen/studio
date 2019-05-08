/*
   @Time : 2019-05-08 17:13
   @Author : frozenchen
   @File : main
   @Software: studio
*/
package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/cmd"
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

	app.Action = func(context *cli.Context) {
		//cli.ShowAppHelp(context)

	}

	app.Commands = append(app.Commands, cli.Command{

		Name: "test",
		Action: func(context *cli.Context) {

			//context.
			fmt.Println("do test")
		},
	})

	cmd.Init(
		cmd.Name("studio"),
		cmd.Description(description),
		cmd.Version(version))

}

func setup() {

}
