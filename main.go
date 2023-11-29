package main

import (
	"os"

	"github.com/ktdocker/commands"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const usage = `ktdocker is a simple container implementation, imitating the docker. 
Available features are: 
	init a container
	run a container
	execute a container
	stop a container
	remove a container
Not implemented many features yet, including container network...`

func main() {
	app := cli.NewApp()
	app.Usage = usage
	app.Name = "ktdocker"

	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}

	app.Commands = []cli.Command{
		commands.InitCommand,
		commands.RunCommand,
		commands.ExecCommand,
		commands.StopCommand,
		commands.RemoveCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
