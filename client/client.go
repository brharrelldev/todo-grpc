package main

import (
	"fmt"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type Tasks struct {
	Tasks []Task `yaml:"tasks"`
}

type Task struct {
	TaskName string `yaml:"task_name"`
}

var client *grpc.ClientConn

func main() {
	app := cli.NewApp()
	app.Name = "todo-client"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "host",
			Usage:    "todo-client --host",
			EnvVar:   "TODO_SERVER_HOST",
			Required: true,
		},
		cli.StringFlag{
			Name:     "port",
			Usage:    "todo-client --port",
			EnvVar:   "TODO_SERVER_PORT",
			Required: true,
		},
	}
	app.Before = grpcPreFlight
	app.Commands = []cli.Command{
		AddCmd,
		ListCmd,
		DeleteCmd,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)

	}

}

func grpcPreFlight(c *cli.Context) error {
	var err error
	client, err = grpc.Dial(net.JoinHostPort(c.String("host"), c.String("port")),
		grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("error dialing grpc server %v", err)
	}

	return nil

}
