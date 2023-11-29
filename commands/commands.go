package commands

import (
	"fmt"
	"os"

	"github.com/ktdocker/cgroups/subsystems"
	"github.com/ktdocker/commands/actions"
	"github.com/ktdocker/container"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var RunCommand = cli.Command{
	Name:  "run",
	Usage: "Create a container with namespace and cgroups limit (e.g., ktdocker run -ti [image] [command])",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		&cli.BoolFlag{
			Name:  "d",
			Usage: "detach container",
		},
		&cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		&cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},
		&cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
		&cli.StringFlag{
			Name:  "name",
			Usage: "container name",
		},
		&cli.StringFlag{
			Name:  "v",
			Usage: "volume",
		},
		&cli.StringSliceFlag{
			Name:  "e",
			Usage: "set environment variables",
		},
	},
	Action: func(context *cli.Context) error {
		if context.NArg() < 1 {
			return fmt.Errorf("missing container command")
		}

		var cmdArray []string
		for _, arg := range context.Args() {
			cmdArray = append(cmdArray, arg)
		}

		imageName := cmdArray[0]
		cmdArray = cmdArray[1:]

		createTty := context.Bool("ti")
		detach := context.Bool("d")
		if createTty && detach {
			return fmt.Errorf("ti and d parameters cannot both be provided")
		}

		resConf := &subsystems.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuSet:      context.String("cpuset"),
			CpuShare:    context.String("cpushare"),
		}

		containerName := context.String("name")
		volume := context.String("v")
		envSlice := context.StringSlice("e")

		actions.RunContainer(createTty, cmdArray, resConf, containerName, volume, imageName, envSlice)
		return nil
	},
}

var InitCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		log.Infof("init command is running")
		return container.RunContainerInitProcess()
	},
}

var ExecCommand = cli.Command{
	Name:  "exec",
	Usage: "Executes a command in a container",
	Action: func(context *cli.Context) error {
		if os.Getenv(actions.ENV_EXEC_PID) != "" {
			log.Infof("Executing PID callback, PID %s", os.Getpid())
			return nil
		}

		if context.NArg() < 2 {
			return fmt.Errorf("missing container name or command")
		}

		containerName := context.Args().Get(0)
		commandArray := context.Args().Tail()

		actions.ExecContainer(containerName, commandArray)
		return nil
	},
}

var StopCommand = cli.Command{
	Name:  "stop",
	Usage: "Stops a container",
	Action: func(context *cli.Context) error {
		if context.NArg() < 1 {
			return fmt.Errorf("missing container name")
		}

		containerName := context.Args().Get(0)
		actions.StopContainer(containerName)
		return nil
	},
}

var RemoveCommand = cli.Command{
	Name:  "rm",
	Usage: "Removes unused containers",
	Action: func(context *cli.Context) error {
		if context.NArg() < 1 {
			return fmt.Errorf("missing container name")
		}

		containerName := context.Args().Get(0)
		actions.RemoveContainer(containerName)
		return nil
	},
}

var LogCommand = cli.Command{
	Name:  "logs",
	Usage: "print logs of a container",
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Please input your container name")
		}
		containerName := context.Args().Get(0)
		actions.LogContainer(containerName)
		return nil
	},
}
